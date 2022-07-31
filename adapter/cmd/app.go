package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"mkstacks/adapter/cli"
	"mkstacks/adapter/gateway"
	"mkstacks/usecase"
)

var appCmd = &cobra.Command{
	Use: "app",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log.Println("init app")
		app := usecase.NewApplication(
			cli.Docker{},
			cli.DockerCompose{},
			gateway.LocalFS{},
		)
		if err := app.CheckDependency(); err != nil {
			return err
		}
		p := context.Background()
		ctx := context.WithValue(p, "App", app)
		cmd.SetContext(ctx)
		return nil
	},
}

var upCmd = &cobra.Command{
	Use: "up",
	RunE: func(cmd *cobra.Command, args []string) error {
		v := cmd.Context().Value("App")
		app, ok := v.(*usecase.Application)
		if !ok {
			return fmt.Errorf("failed to init application")
		}
		return app.Up(args)
	},
}

var downCmd = &cobra.Command{
	Use: "down",
	RunE: func(cmd *cobra.Command, args []string) error {
		v := cmd.Context().Value("App")
		app, ok := v.(*usecase.Application)
		if !ok {
			return fmt.Errorf("failed to init application")
		}
		return app.Down()
	},
}

var recreateCmd = &cobra.Command{
	Use: "recreate",
	RunE: func(cmd *cobra.Command, args []string) error {
		v := cmd.Context().Value("App")
		app, ok := v.(*usecase.Application)
		if !ok {
			return fmt.Errorf("failed to init application")
		}
		return app.Recreate(args)
	},
}

var buildCmd = &cobra.Command{
	Use: "build",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	appCmd.AddCommand(upCmd)
	appCmd.AddCommand(downCmd)
	appCmd.AddCommand(recreateCmd)
	RootCmd.AddCommand(appCmd)
}

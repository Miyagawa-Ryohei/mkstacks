package cmd

import (
	"github.com/spf13/cobra"
	"mkstacks/adapter/gateway"
	"mkstacks/usecase"
)

var InitCmd = &cobra.Command{
	Use: "init",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializer := usecase.NewInitializer(gateway.LocalFS{})
		return initializer.Init()
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
}

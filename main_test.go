package main_test

import (
	"mkstacks/adapter/cli"
	"mkstacks/adapter/gateway"
	"mkstacks/usecase"
	"testing"
)

func TestApplicationOnSuccess(t *testing.T) {
	app := usecase.NewApplication(cli.Docker{}, cli.DockerCompose{}, gateway.LocalFS{})
	t.Run("Up", func(t *testing.T) {
		if err := app.Up([]string{"opensearch"}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Down", func(t *testing.T) {
		if err := app.Down(); err != nil {
			t.Fatal(err)
		}
	})
}

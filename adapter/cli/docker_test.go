package cli_test

import (
	"mkstacks/adapter/cli"
	"testing"
)

func TestDocker_IsInstalled(t *testing.T) {
	t.Run("is docker installed", func(t *testing.T) {
		d := cli.Docker{}
		if ans := d.IsInstalled(); ans != nil {
			t.Fatal("docker is not installed")
		}
	})
}

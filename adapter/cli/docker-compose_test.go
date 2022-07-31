package cli_test

import (
	"mkstacks/adapter/cli"
	"os"
	"testing"
)

func TestDockerCompose(t *testing.T) {
	d := cli.DockerCompose{}
	t.Run("IsInstalled", func(t *testing.T) {
		if ans := d.IsInstalled(); ans != nil {
			t.Fatal("docker is not installed")
		}
	})
	t.Run("Up", func(t *testing.T) {
		composes := []string{
			"/home/ryo/development/Private/mkstacks/docker-compose.yml",
			"/home/ryo/development/Private/mkstacks/overrides/docker-compose.override-opensearch.yml",
		}
		buf, err := d.Up(composes)
		if err != nil {
			t.Fatal(err)
		}
		os.WriteFile("/home/ryo/development/Private/mkstacks/overrides/mkstacks.yml", buf, 0755)
		t.Cleanup(func() {
			d.Down(composes)
		})
	})
	t.Run("docker-compose Down test", func(t *testing.T) {
		composes := []string{
			"/home/ryo/development/Private/mkstacks/docker-compose.yml",
			"/home/ryo/development/Private/mkstacks/overrides/docker-compose.override-opensearch.yml",
		}
		err := d.Down(composes)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.Remove("/home/ryo/development/Private/mkstacks/overrides/mkstacks.yml"); err != nil {
			t.Fatal(err)
		}
	})

}

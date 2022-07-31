package cli

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

type DockerCompose struct{}

func (dc DockerCompose) IsInstalled() (err error) {
	cmd := exec.Command("docker-compose", "version")
	err = cmd.Run()
	return
}

func (dc DockerCompose) Up(composes []string) ([]byte, error) {
	arg := make([]string, 0)
	for _, compose := range composes {
		arg = append(arg, "-f", compose)
	}
	arg = append(arg)
	cmd := exec.Command("docker-compose", append(arg, "up", "-d")...)
	conf := exec.Command("docker-compose", append(arg, "config")...)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	buf, err := conf.Output()
	if err != nil {
		if e := dc.Down(composes); e != nil {
			return nil, errors.Wrap(e, err.Error())
		}
		return nil, err
	}

	return buf, nil
}

func (dc DockerCompose) Down(composes []string) error {
	arg := make([]string, 0)
	for _, compose := range composes {
		arg = append(arg, "-f", compose)
	}
	arg = append(arg)
	cmd := exec.Command("docker-compose", append(arg, "down")...)
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

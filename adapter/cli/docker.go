package cli

import "os/exec"

type Docker struct{}

func (d Docker) IsInstalled() (err error) {
	cmd := exec.Command("docker", "version")
	err = cmd.Run()
	return
}

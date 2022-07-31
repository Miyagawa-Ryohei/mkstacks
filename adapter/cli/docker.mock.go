package cli

import "fmt"

type DockerSuccessMock struct{}

func (d DockerSuccessMock) IsInstalled() (err error) {
	return
}

type DockerErrorMock struct{}

func (d DockerErrorMock) IsInstalled() (err error) {
	return fmt.Errorf("return error")
}

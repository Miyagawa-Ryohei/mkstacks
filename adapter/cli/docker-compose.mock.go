package cli

import (
	"bytes"
	"fmt"
)

type DockerComposeSuccessMock struct{}

func (dc DockerComposeSuccessMock) IsInstalled() (err error) {
	return
}

func (dc DockerComposeSuccessMock) Up(composes []string) ([]byte, error) {
	return bytes.NewBufferString("up is success").Bytes(), nil
}

type DockerComposeErrorMock struct{}

func (dc DockerComposeErrorMock) IsInstalled() (err error) {
	return fmt.Errorf("return error")
}

func (dc DockerComposeErrorMock) Up(composes []string) ([]byte, error) {
	return nil, fmt.Errorf("return error")
}

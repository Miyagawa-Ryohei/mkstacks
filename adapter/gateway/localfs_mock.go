package gateway

import "fmt"

type LocalFSSuccessMock struct {
	ReadResult []byte
}

func (l LocalFSSuccessMock) Write(data interface{}, locate string) error {
	return nil
}

func (l LocalFSSuccessMock) List(locate string, recursive bool) error {
	return nil
}

func (l LocalFSSuccessMock) Read(locate string) ([]byte, error) {
	return l.ReadResult, nil
}

type LocalFSErrorMock struct {
	ReadResult []byte
}

func (l LocalFSErrorMock) Write(data interface{}, locate string) error {
	return fmt.Errorf("dummy write error")
}

func (l LocalFSErrorMock) List(locate string, recursive bool) error {
	return fmt.Errorf("dummy list error")
}

func (l LocalFSErrorMock) Read(locate string) ([]byte, error) {
	return nil, fmt.Errorf("dummy read error")
}

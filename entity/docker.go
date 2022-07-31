package entity

type Docker interface {
	IsInstalled() error
}

type DockerCompose interface {
	IsInstalled() error
	Up(composes []string) ([]byte, error)
	Down(composes []string) error
}

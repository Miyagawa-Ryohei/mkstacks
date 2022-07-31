package entity

type Storage interface {
	CreatDir(path string) error
	Write(data []byte, path string) error
	WriteAsYaml(data interface{}, path string) error
	Read(path string) ([]byte, error)
}

package gateway

import (
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type LocalFS struct {
}

func (l LocalFS) CreatDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func (l LocalFS) prepare(locate string) error {
	target := ""
	d, _ := path.Split(locate)
	if !strings.Contains(locate, string(os.PathSeparator)) {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		target = path.Join(cwd, locate)
		goto FileCreate
	}
	target = locate
	if err := os.MkdirAll(d, 0755); err != nil {
		return err
	}
	if _, err := os.Stat(locate); err == nil {
		return nil
	}
FileCreate:
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	return nil
}

func (l LocalFS) Write(data []byte, locate string) error {
	if err := l.prepare(locate); err != nil {
		return err
	}
	if err := os.WriteFile(locate, data, fs.ModeAppend); err != nil {
		return err
	}
	return nil
}
func (l LocalFS) WriteAsYaml(data interface{}, locate string) error {
	b, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return l.Write(b, locate)
}

func (l LocalFS) List(locate string, recursive bool) error {
	return nil
}

func (l LocalFS) Read(locate string) ([]byte, error) {
	f, err := os.Open(locate)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

package usecase

import (
	"bytes"
	"fmt"
	"mkstacks/entity"
	"os"
	"path"
)

type Initializer struct {
	filer entity.Storage
	cwd   string
}

func (i *Initializer) Init() error {
	dotStack := fmt.Sprintf("%s", `
## user definition override file setting:
# udofs:
#  opensearch:
#    depends:
#      - mine

## environment of every service common setting
# parameters:
#  DUMMY_PARAMETER1: AAA
#  DUMMY_PARAMETER2: AAA
#  DUMMY_PARAMETER3: AAA

## extra_hosts of every service common setting
# hosts:
#  DUMMY_HOST_A: 192.168.11.1
#  DUMMY_HOST_B: 192.168.11.1
#  DUMMY_HOST_C: 192.168.11.1

`)
	if err := i.filer.Write(bytes.NewBufferString(dotStack).Bytes(), path.Join(i.cwd, ".stack.yml")); err != nil {
		return err
	}
	if err := i.filer.CreatDir(path.Join(i.cwd, "overrides")); err != nil {
		return err
	}
	return nil
}

func NewInitializer(filer entity.Storage) *Initializer {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &Initializer{
		filer,
		cwd,
	}
}

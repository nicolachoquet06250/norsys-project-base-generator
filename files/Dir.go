package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	. "test_go_webserver/helpers"
)

type (
	FileSystemDirectory interface {
		Exists() (exists bool, err error)
		Create() error
		Is() bool
	}

	Dir struct {
		Path string

		FileSystemDirectory
	}
)

func (d Dir) Exists() (exists bool, err error) {
	d.Path = strings.ReplaceAll(d.Path, Slash()+Slash(), Slash())

	if d.Is() {
		_, err := ioutil.ReadDir(d.Path)
		if err != nil {
			return false, nil
		}
		return true, nil
	}

	return false, fmt.Errorf("le répertoire demandé n'existe pas")
}

func (d Dir) Create() error {
	d.Path = strings.ReplaceAll(d.Path, Slash()+Slash(), Slash())

	var splitPath = strings.Split(d.Path, Slash())

	_, _ = ArrayPop(&splitPath)

	p := strings.Join(splitPath, Slash())
	err := os.MkdirAll(p, 0777)
	if err != nil {
		return fmt.Errorf("%s directory create error", err.Error())
	}

	return nil
}

func (d Dir) Is() bool {
	if !strings.Contains(d.Path, ".") {
		return true
	}
	return false
}

func NewDir(path string) Dir {
	return Dir{
		Path: path,
	}
}

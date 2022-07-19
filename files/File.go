package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	. "test_go_webserver/helpers"
)

type FileSystemFile interface {
	Exists(path string) (exists bool, err error)
	Create(path string, content string, recursive bool) error
	Is(path string) bool
}

type File struct {
	Path string

	FileSystemFile
}

func (f File) Exists() (exists bool, err error) {
	f.Path = strings.ReplaceAll(f.Path, Slash()+Slash(), Slash())

	if f.Is() {
		_, err := ioutil.ReadFile(f.Path)
		if err != nil {
			return false, nil
		}
		return true, nil
	}

	return false, fmt.Errorf("le fichier demandé n'existe pas")
}

func (f File) Create(content string, recursive bool) error {
	f.Path = strings.ReplaceAll(f.Path, Slash()+Slash(), Slash())

	var splitPath = strings.Split(f.Path, Slash())

	_, _ = ArrayPop(&splitPath)

	p := strings.Join(splitPath, Slash())

	if recursive {
		err := os.MkdirAll(p, 0777)
		if err != nil {
			return fmt.Errorf("%s directory create error", err.Error())
		}
	} else {
		exists, _ := NewDir(p).Exists()
		if !exists {
			return fmt.Errorf("%s directory not exists", p)
		}
	}

	exists, _ := f.Exists()
	if exists {
		return fmt.Errorf("%s file already exists", f.Path)
	}

	_f, err := os.Create(f.Path)
	if err != nil {
		return fmt.Errorf("%s file open error", err.Error())
	}

	_, err = _f.Write([]byte(content))
	if err != nil {
		return fmt.Errorf("%s file write error", err.Error())
	}

	err = _f.Close()
	if err != nil {
		return fmt.Errorf("%s file close error", err.Error())
	}

	return nil
}

func (f File) Is() bool {
	if strings.Contains(f.Path, ".") {
		return true
	}
	return false
}

func NewFile(path string) File {
	return File{
		Path: path,
	}
}

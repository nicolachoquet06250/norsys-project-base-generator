package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	. "test_go_webserver/helpers"
)

type (
	FileSystemFile interface {
		Exists() (exists bool, err error)
		Create(content string, recursive bool) error
		Is() bool
		GetContent() (content string, err error)
		Update(content string) error
	}

	File struct {
		Path string
	}
)

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

func (f File) GetContent() (content string, err error) {
	c, e := os.ReadFile(f.Path)
	return string(c), e
}

func (f File) Update(content string) error {
	// Read Write Mode
	file, err := os.OpenFile(f.Path, os.O_RDWR, 0644)

	if err != nil {
		return fmt.Errorf("failed opening file: %s", err)
	}
	defer file.Close()

	_, err = file.WriteAt([]byte(content), 0) // Write at 0 beginning
	if err != nil {
		return fmt.Errorf("failed writing to file: %s", err)
	}
	//fmt.Printf("\nLength: %d bytes", l)
	//fmt.Printf("\nFile Name: %s", file.Name())
	return nil
}

func NewFile(path string) File {
	return File{
		Path: path,
	}
}

package files

import (
	"fmt"
	"strings"
	"test_go_webserver/configFiles"
	. "test_go_webserver/helpers"
	. "test_go_webserver/pages/helpers"
	"test_go_webserver/technos"
)

type (
	FileSystem interface {
		Create(techno technos.Techno) (alert Alert)
		Exists() (exists bool, err error)
	}

	Project struct {
		Path string `json:"path"`
		Name string `json:"name"`
	}
)

func (p Project) Create(techno technos.Techno) (alert Alert) {
	var (
		err    error
		exists bool
	)

	technoName := techno.Name
	technoValue := techno.Value
	completePath := p.Path + Slash() + p.Name

	exists, err = p.Exists()
	if err == nil && exists {
		alert = Alert{
			Message: fmt.Sprintf("Le projet %s existe déjà dans le répertoire %s !", p.Name, p.Path),
			Type:    ERROR,
		}
		println("Le projet " + p.Name + " existe déjà dans le répertoire " + p.Path + " !")

		return alert
	}

	for _p, content := range configFiles.ConfigFiles[technoName] {
		params := map[string]interface{}{"ProjectName": &p.Name}

		_p, err = ParseString("path", _p, params)
		if err != nil {
			return Alert{
				Message: fmt.Sprintf("erreur template path: %s", err.Error()),
				Type:    ERROR,
			}
		}

		content, err = ParseString("content", content, params)
		if err != nil {
			return Alert{
				Message: fmt.Sprintf("erreur template content: %s", err.Error()),
				Type:    ERROR,
			}
		}

		err = NewFile(completePath+Slash()+_p).Create(content, true)
		if err != nil {
			return Alert{
				Message: fmt.Sprintf("erreur: %s", err.Error()),
				Type:    ERROR,
			}
		}
	}

	exists, err = p.Exists()
	if err == nil && exists {
		alert = Alert{
			Message: fmt.Sprintf("Le projet %s à bien été généré dans le répertoire %s !", technoName, completePath),
			Type:    SUCCESS,
		}
	} else {
		alert = Alert{
			Message: fmt.Sprintf("Une erreur est survenue lors de la génération du projet %s dans le répertoire %s !", technoValue, p.Path),
			Type:    ERROR,
		}
	}

	return alert
}

func (p Project) Exists() (exists bool, err error) {
	return NewDir(p.Path + Slash() + p.Name).Exists()
}

func NewProject(path string, name *string) Project {
	if name == nil {
		splitted := strings.Split(path, Slash())
		name, _ = ArrayPop(&splitted)
		path = strings.Join(splitted, Slash())
	}

	return Project{
		Path: path,
		Name: *name,
	}
}

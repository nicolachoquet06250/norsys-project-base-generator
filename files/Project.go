package files

import (
	"fmt"
	"npbg/configFiles"
	. "npbg/helpers"
	. "npbg/pages/helpers"
	"npbg/technos"
	"strings"
)

type (
	FileSystem interface {
		Create() (alert Alert)
		Remove(id int) (alert Alert)
		Exists() (exists bool, err error)
	}

	Project struct {
		Path   string         `json:"path"`
		Name   string         `json:"name"`
		Techno technos.Techno `json:"techno"`
	}
)

func (p Project) Create() (alert Alert) {
	var (
		err    error
		exists bool
	)

	technoName := p.Techno.Name
	technoValue := p.Techno.Value
	completePath := p.Path + Slash() + p.Name

	exists, err = p.Exists()
	if err == nil && exists {
		alert := NewAlert(fmt.Sprintf("Le projet %s existe déjà dans le répertoire %s !", p.Name, p.Path), ERROR)
		println("Le projet " + p.Name + " existe déjà dans le répertoire " + p.Path + " !")

		return alert
	}

	for _p, content := range configFiles.ConfigFiles[technoName] {
		params := map[string]interface{}{"ProjectName": &p.Name}

		_p, err = ParseString("path", _p, &params)
		if err != nil {
			return NewAlert(fmt.Sprintf("erreur template path: %s", err.Error()), ERROR)
		}

		content, err = ParseString("content", content, &params)
		if err != nil {
			return NewAlert(fmt.Sprintf("erreur template content: %s", err.Error()), ERROR)
		}

		err = NewFile(completePath+Slash()+_p).Create(content, true)
		if err != nil {
			return NewAlert(fmt.Sprintf("erreur: %s", err.Error()), ERROR)
		}
	}

	exists, err = p.Exists()
	if err == nil && exists {
		alert = NewAlert(
			fmt.Sprintf("Le projet %s à bien été généré dans le répertoire %s !", technoName, completePath),
			SUCCESS,
		)
	} else {
		alert = NewAlert(
			fmt.Sprintf("Une erreur est survenue lors de la génération du projet %s dans le répertoire %s !", technoValue, p.Path),
			ERROR,
		)
	}

	return alert
}

func (p Project) Remove() (alert Alert) {
	err := NewDir(p.Path + Slash() + p.Name).Remove()

	if err != nil {
		return NewAlert(err.Error(), ERROR)
	}

	return NewAlert("Le projet as été supprimé avec succès", SUCCESS)
}

func (p Project) Exists() (exists bool, err error) {
	return NewDir(p.Path + Slash() + p.Name).Exists()
}

func NewProject(path string, name *string, techno technos.Techno) Project {
	if name == nil {
		splitted := strings.Split(path, Slash())
		name, _ = ArrayPop(&splitted)
		path = strings.Join(splitted, Slash())
	}

	return Project{
		Path:   path,
		Name:   *name,
		Techno: techno,
	}
}

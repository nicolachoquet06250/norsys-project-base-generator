package history

import (
	"encoding/json"
	"fmt"
	"npbg/files"
	"npbg/helpers"
	"npbg/technos"
	"os/user"
	"strings"
)

type (
	IHistory interface {
		AddProject() error
		RemoveProject() error
		UpdateProject(id int) error
	}

	ItemHistory files.Project

	History []ItemHistory
)

var history = &History{}

func GetProjectArchivesPath() string {
	var (
		username = func() string {
			u, _ := user.Current()
			return u.Username
		}()
		historyFilePath string
	)

	if helpers.IsWindows() {
		historyFilePath = helpers.RootPath() + "npbg"
	} else {
		historyFilePath = helpers.RootPath() + "home" + helpers.Slash() + username + helpers.Slash() + "npbg"
	}

	return historyFilePath
}

func GetIconPath() string {
	return GetProjectArchivesPath() + helpers.Slash() + "logo-norsys.png"
}

func GetHistoryFilePath() string {
	return GetProjectArchivesPath() + helpers.Slash() + "npbg-history.json"
}

func (h History) Add(project ItemHistory) (alert helpers.Alert) {
	err := project.AddProject()

	if err != nil {
		return helpers.NewAlert(err.Error(), helpers.ERROR)
	}

	return helpers.NewAlert("Le project as été ajouté avec succès à l'historique", helpers.SUCCESS)
}

func (h History) Remove(project ItemHistory) (alert helpers.Alert) {
	err := project.RemoveProject()

	if err != nil {
		return helpers.NewAlert(err.Error(), helpers.ERROR)
	}

	return helpers.NewAlert("Le project as été supprimé avec succès de l'historique", helpers.SUCCESS)
}

func (h History) Exists(project ItemHistory) bool {
	_h := GetHistoryList()

	nh := helpers.ArrayFilter(*_h, func(e ItemHistory, i int) bool {
		if e.Path == project.Path && e.Name == project.Name && e.Techno.Name == project.Techno.Name {
			return true
		}
		return false
	})

	return len(nh) > 0
}

func (h History) IsEmpty() bool {
	return len(h) == 0
}

func (h ItemHistory) AddProject() error {
	hist := GetHistoryList()
	var oldCmp = len(*hist)
	*hist = append(*hist, h)
	var newCmp = len(*hist)
	if oldCmp == newCmp {
		return fmt.Errorf("impossible d'ajouter le projet %s à l'historique", h.Name)
	}

	jsonV, err := json.Marshal(*hist)
	if err != nil {
		return err
	}

	file := files.NewFile(GetHistoryFilePath())

	err = file.Update(string(jsonV))
	if err != nil {
		return err
	}

	return nil
}

func (h ItemHistory) RemoveProject() error {
	hist := GetHistoryList()

	oldSize := len(*hist)

	nh := helpers.ArrayFilter(*hist, func(e ItemHistory, i int) bool {
		if e.Path != h.Path || e.Name != h.Name || e.Techno.Name != h.Techno.Name {
			return true
		}
		return false
	})

	newSize := len(nh)

	if oldSize <= newSize {
		return fmt.Errorf("une erreur est survenue lors de la suppression du projet de l'historique")
	}

	jsonV, err := json.Marshal(nh)
	if err != nil {
		return err
	}

	file := files.NewFile(GetHistoryFilePath())

	_ = file.Empty()

	err = file.Update(string(jsonV))
	if err != nil {
		return err
	}

	return nil
}

func (h ItemHistory) UpdateProject(id int) error {
	return nil
}

func GetHistoryList() *History {
	file := files.NewFile(GetHistoryFilePath())

	historyFileContent, err := file.GetContent()

	if err != nil {
		println(err.Error())
		historyFileContent = "[]"
		_ = file.Create(historyFileContent, true)
	}

	historyFileContent, err = file.GetContent()

	if err == nil {
		err = json.Unmarshal([]byte(historyFileContent), history)

		if err != nil {
			println(err.Error())
		}
	} else {
		println(err.Error())
	}

	return history
}

func NewItem(path string, name *string, techno technos.Techno) ItemHistory {
	if name == nil {
		splitted := strings.Split(path, helpers.Slash())
		name, _ = helpers.ArrayPop(&splitted)
		path = strings.Join(splitted, helpers.Slash())
	}

	return ItemHistory{
		Path:   path,
		Name:   *name,
		Techno: techno,
	}
}

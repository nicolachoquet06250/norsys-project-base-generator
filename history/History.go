package history

import (
	"encoding/json"
	"fmt"
	"os/user"
	"strings"
	"test_go_webserver/files"
	"test_go_webserver/helpers"
	"test_go_webserver/technos"
)

type (
	IHistory interface {
		AddProject() error
		RemoveProject(id int) error
		UpdateProject(id int) error
	}

	ItemHistory files.Project

	History []ItemHistory
)

var history = &History{}

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

	var (
		username = func() string {
			u, _ := user.Current()
			return u.Username
		}()
		historyFileName = "npbg-history.json"
		historyFilePath string
	)

	if helpers.IsWindows() {
		historyFilePath = helpers.RootPath() + "npbg"
	} else {
		historyFilePath = helpers.RootPath() + "home" + helpers.Slash() + username + helpers.Slash() + "npbg"
	}

	file := files.NewFile(historyFilePath + helpers.Slash() + historyFileName)

	err = file.Update(string(jsonV))
	if err != nil {
		return err
	}

	return nil
}

func (h ItemHistory) RemoveProject(id int) error {
	return nil
}

func (h ItemHistory) UpdateProject(id int) error {
	return nil
}

func GetHistoryList() *History {
	var (
		username string = func() string {
			u, _ := user.Current()
			return u.Username
		}()
		historyFileName string = "npbg-history.json"
		historyFilePath string
	)

	if helpers.IsWindows() {
		historyFilePath = helpers.RootPath() + "npbg"
	} else {
		historyFilePath = helpers.RootPath() + "home" + helpers.Slash() + username + helpers.Slash() + "npbg"
	}

	file := files.NewFile(historyFilePath + helpers.Slash() + historyFileName)

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

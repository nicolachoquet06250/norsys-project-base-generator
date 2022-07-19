package history

import (
	"encoding/json"
	"fmt"
	"os/user"
	"test_go_webserver/files"
	"test_go_webserver/helpers"
)

type (
	IHistory interface {
		AddProject() error
		RemoveProject(id int) error
		UpdateProject(id int) error
	}

	History struct {
		project files.Project

		IHistory
	}
)

var historyList []History

func NewHistory(project files.Project) History {
	username := func() string {
		u, _ := user.Current()
		return u.Username
	}()

	historyFileName := "npbg-history.json"
	var historyFilePath string
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

	println(historyFileContent)

	err = json.Unmarshal([]byte(historyFileContent), &historyList)

	fmt.Printf("History : %+v, %d \n", &historyList, len(historyList))

	if err != nil {
		println(err.Error())
	}

	for _, history := range historyList {
		if history.project.Name == nil {
			println("Path : ", history.project.Path)
		} else {
			println(
				"Path : ", history.project.Path,
				"Name : ", *history.project.Name,
			)
		}
	}

	d, _ := json.Marshal(historyList)

	println(d)

	return History{project: project}
}

func (h History) AddProject() error {
	return nil
}

func (h History) RemoveProject(id int) error {
	return nil
}

func (h History) UpdateProject(id int) error {
	return nil
}

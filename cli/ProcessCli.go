package cli

import (
	"fmt"
	"os"
	"test_go_webserver/files"
	. "test_go_webserver/helpers"
	"test_go_webserver/history"
	"test_go_webserver/technos"
)

func ProcessCli() (exit bool) {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		if args[0] == "help" {
			Help()
			return true
		}

		projectPath := args[0]
		techno, technoErr := technos.FromValue(args[1])

		if technoErr != nil {
			println(technoErr.Error())
			return true
		}

		project := files.NewProject(projectPath, nil, techno)

		var alert Alert
		exists, _ := project.Exists()
		if exists {
			alert = NewAlert(fmt.Sprintf("Le projet %s existe déjà dans le répertoire %s !", project.Name, project.Path), ERROR)
		} else {
			alert = project.Create()

			if alert.Type == SUCCESS {
				item := history.NewItem(project.Path, &project.Name, project.Techno)
				err := item.AddProject()
				if err != nil {
					alert = NewAlert(err.Error(), ERROR)
				}
			}
		}

		alertMessage := ""
		if alert.Type == ERROR {
			alertMessage += "ERROR: "
		}

		alertMessage += alert.Message

		println(alertMessage)

		return true
	}

	return false
}

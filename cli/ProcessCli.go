package cli

import (
	"fmt"
	"npbg/files"
	. "npbg/helpers"
	"npbg/history"
	"npbg/technos"
	"os"
)

func Process() (exit bool) {
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

			if alert.Type == string(SUCCESS) {
				item := history.NewItem(project.Path, &project.Name, project.Techno)
				err := item.AddProject()
				_alert := MaybeError(err, func(err error) *Alert {
					r := NewAlert(err.Error(), ERROR)
					return &r
				})
				if _alert != nil {
					alert = *_alert
				}
			}
		}

		alertMessage := ""
		if alert.Type == string(ERROR) {
			alertMessage += "ERROR: "
		}

		alertMessage += alert.Message

		println(alertMessage)

		return true
	}

	return false
}

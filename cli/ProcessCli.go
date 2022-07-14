package cli

import (
	"os"
	"test_go_webserver/files"
	"test_go_webserver/helpers"
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

		alert := files.ProjectGeneration(projectPath, techno)

		alertMessage := ""
		if alert.Type == helpers.ERROR {
			alertMessage += "ERROR: "
		}

		alertMessage += alert.Message

		println(alertMessage)

		return true
	}

	return false
}

package main

import (
	"log"
	"net/http"
	"os"
	"test_go_webserver/cli"
	"test_go_webserver/files"
	"test_go_webserver/http/routing"
	"test_go_webserver/technos"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		if args[0] == "help" {
			cli.Help()
			return
		}

		projectPath := args[0]
		techno, technoErr := technos.FromValue(args[1])

		if technoErr != nil {
			println(technoErr.Error())
			return
		}

		files.ProjectGeneration(projectPath, techno)
		return
	}

	routing.Routes()

	println("server opened on http://localhost:8042")
	cli.Browser{}.Open("http://localhost:8042")
	err := http.ListenAndServe(":8042", nil)
	if err != nil {
		log.Fatal(err)
	}
}

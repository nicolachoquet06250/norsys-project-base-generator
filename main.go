package main

import (
	"log"
	"net/http"
	"os"
	"test_go_webserver/cli"
	"test_go_webserver/files"
	"test_go_webserver/http/httpMethods"
	. "test_go_webserver/pages"
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

	httpMethods.HandleGet("/", Home)
	httpMethods.HandleGet("/generate", Generate)
	httpMethods.HandleGet("/help", Help)

	println("server opened on http://localhost:8042")
	log.Fatal(http.ListenAndServe(":8042", nil))
}

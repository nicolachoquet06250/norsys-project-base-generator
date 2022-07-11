package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"test_go_webserver/cli"
	"test_go_webserver/files"
	"test_go_webserver/http/portChoice"
	"test_go_webserver/http/routing"
	"test_go_webserver/technos"
)

func Process(registerRoutes bool) {
	port := portChoice.ChooseUnusedPort()

	if registerRoutes == true {
		routing.Routes()
	}

	portStr := strconv.FormatInt(int64(port), 10)
	portSuffix := ":" + portStr

	if registerRoutes == true {
		cli.Browser{}.Open("http://localhost" + portSuffix)
		println("server opened on http://localhost" + portSuffix)
	}

	err := http.ListenAndServe(portSuffix, nil)

	if err != nil {
		if strings.Contains(err.Error(), "invalid port") == true {
			println(err.Error())
			Process(false)
		} else {
			log.Fatal(err)
		}
	}
}

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

	Process(true)
}

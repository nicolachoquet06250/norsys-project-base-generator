package main

import (
	"test_go_webserver/cli"
	"test_go_webserver/server"
)

func main() {
	exit := cli.ProcessCli()

	if exit {
		return
	}

	//projectName := "php-project"
	//history.NewHistory(files.NewProject("C:\\Users\\NDZQ5522\\Desktop", &projectName))

	server.Process(true, true)
}

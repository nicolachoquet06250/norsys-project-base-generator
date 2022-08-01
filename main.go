package main

import (
	"os"
	"test_go_webserver/cli"
	"test_go_webserver/server"
)

func main() {
	exit := cli.ProcessCli()

	if exit {
		return
	}

	server.Process(true, os.Getenv("OPEN_BROWSER") == "1" || os.Getenv("OPEN_BROWSER") == "")
}

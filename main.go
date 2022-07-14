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

	server.Process(true, true)
}

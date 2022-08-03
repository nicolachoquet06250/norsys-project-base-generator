package main

import (
	"npbg/cli"
	"npbg/helpers"
	"npbg/server"
	"os"
	"strings"
)

func main() {
	exit := cli.Process()

	if exit {
		return
	}

	var openBrowser = os.Getenv("OPEN_BROWSER") == "1" || os.Getenv("OPEN_BROWSER") == "" ||
		!strings.Contains(os.Args[0], helpers.Slash()+"b001"+helpers.Slash()+"exe"+helpers.Slash())

	server.Process(true, openBrowser)
}

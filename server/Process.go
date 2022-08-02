package server

import (
	"log"
	"net/http"
	"npbg/cli"
	"npbg/http/portChoice"
	"npbg/http/routing"
	"strconv"
	"strings"
)

func Process(registerRoutes bool, openBrowser bool) {
	port := portChoice.ChooseUnusedPort()

	if registerRoutes == true {
		routing.Routes()
	}

	portStr := strconv.FormatInt(int64(port), 10)
	portSuffix := ":" + portStr

	if registerRoutes == true {
		println("server opened on http://localhost" + portSuffix)
	}

	if openBrowser == true {
		cli.Browser{}.Open("http://localhost" + portSuffix)
	}

	err := http.ListenAndServe(portSuffix, nil)

	if err != nil {
		if strings.Contains(err.Error(), "invalid port") == true {
			println(err.Error())
			Process(false, openBrowser)
		} else {
			log.Fatal(err)
		}
	}
}

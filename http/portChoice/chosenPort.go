package portChoice

import (
	"regexp"
	"strconv"
	"strings"
	"test_go_webserver/cli"
	. "test_go_webserver/helpers"
)

var ChosenPort int

func ChooseUnusedPort() int {
	out := cli.ExeCmd("netstat -nao")

	lines := strings.Split(string(out), cli.BackLine())

	var listeningPorts []int

	for _, l := range lines {
		var re = regexp.MustCompile(`^ {2}TCP {4}\d{0,3}\.\d{0,3}\.\d{0,3}\.\d{0,3}:(?P<local_port>80\d+) +\d{0,3}\.\d{0,3}\.\d{0,3}\.\d{0,3}:(?P<remote_port>\d+) +(LISTENING) +\d+$`)

		for _, match := range re.FindAllString(l, -1) {
			matches := re.FindStringSubmatch(match)

			localPort, _ := strconv.Atoi(matches[re.SubexpIndex("local_port")])

			listeningPorts = append(listeningPorts, localPort)
		}
	}

	localPort := RandomNumber(8000, 8099)

	chosenPort := localPort

	exists, _ := InArray(chosenPort, listeningPorts)
	if exists == true {
		return ChooseUnusedPort()
	} else {
		ChosenPort = localPort
		return ChosenPort
	}
}

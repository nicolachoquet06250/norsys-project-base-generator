package portChoice

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"test_go_webserver/cli"
	"test_go_webserver/helpers"
)

var ChosenPort int

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func ChooseUnusedPort() int {
	out := cli.ExeCmd("netstat -nao")

	lines := strings.Split(string(out), cli.BackLine())

	var listeningPorts []int

	for _, l := range lines {
		var re = regexp.MustCompile(`^ {2}TCP {4}\d{0,3}\.\d{0,3}\.\d{0,3}\.\d{0,3}:(?P<local_port>80\d+) +\d{0,3}\.\d{0,3}\.\d{0,3}\.\d{0,3}:(?P<remote_port>\d+) +(LISTENING) +\d+$`)

		for _, match := range re.FindAllString(l, -1) {
			println(match)
			matches := re.FindStringSubmatch(match)

			localPort, _ := strconv.Atoi(matches[re.SubexpIndex("local_port")])

			listeningPorts = append(listeningPorts, localPort)
		}
	}

	localPort := helpers.RandomNumber(8000, 8099)

	chosenPort := localPort

	exists, _ := InArray(chosenPort, listeningPorts)
	println(len(listeningPorts))
	for _, p := range listeningPorts {
		println(p)
	}
	println(strconv.FormatInt(int64(chosenPort), 10)+" exists", exists)
	if exists == true {
		return ChooseUnusedPort()
	} else {
		ChosenPort = localPort
		return ChosenPort
	}
}

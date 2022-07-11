package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func BackLine() string {
	if os.Getenv("GOOS") == "windows" {
		return "\r\n"
	} else if os.Getenv("GOOS") == "linux" || os.Getenv("GOOS") == "darwin" {
		return "\n"
	}

	return ""
}

func ExeCmd(cmd string) []byte {
	parts := strings.Fields(cmd)

	out, err := exec.Command(parts[0], parts[1]).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}

	return out
}

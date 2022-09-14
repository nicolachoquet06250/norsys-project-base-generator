package cli

import (
	"fmt"
	. "npbg/helpers"
	"os/exec"
	"runtime"
	"strings"
)

func BackLine() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		return "\n"
	}

	return ""
}

func ExeCmd(cmd string) []byte {
	parts := strings.Fields(cmd)

	out, err := exec.Command(parts[0], parts[1]).Output()
	MaybeError(err, func(err error) *int64 {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
		var r *int64
		return r
	})

	return out
}

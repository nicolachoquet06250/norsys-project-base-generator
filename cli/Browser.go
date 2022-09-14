package cli

import (
	"fmt"
	"log"
	. "npbg/helpers"
	"os/exec"
	"runtime"
)

type Browser struct{}

func (_ Browser) Open(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	MaybeError(err, func(err error) *int64 {
		log.Fatal(err)
		var r *int64
		return r
	})
}

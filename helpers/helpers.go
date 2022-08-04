package helpers

import (
	"os"
	"os/user"
	"runtime"
	"strings"
)

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}

func IsDarwin() bool {
	return runtime.GOOS == "darwin"
}

func CharBack() string {
	return "\n"
}

func Tab() string {
	return "\t"
}

func Slash() string {
	var slash = "/"

	if IsWindows() {
		slash = "\\"
	}

	return slash
}

func PwdVar() string {
	start := "("
	end := ")"

	if IsWindows() {
		start = "{"
		end = "}"
	}

	return "$" + start + "pwd" + end
}

func RootPath() string {
	if IsWindows() {
		return "C:\\"
	}
	return "/"
}

func HomePath() string {
	path := RootPath()
	username := func() string {
		u, _ := user.Current()
		return u.Username
	}()

	if IsWindows() {
		path += "Users" + Slash() + strings.Split(username, Slash())[1]
	} else {
		path += "home" + Slash() + username
	}

	return path
}

func IsBuild() bool {
	return os.Getenv("GOBUILD") != "1"
}

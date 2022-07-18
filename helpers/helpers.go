package helpers

import (
	"os"
	"runtime"
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
	var root = "/"

	if IsWindows() {
		root = "C:\\"
	}

	return root
}

func IsBuild() bool {
	return os.Getenv("GOBUILD") != "1"
}

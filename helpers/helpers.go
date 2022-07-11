package helpers

import (
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"test_go_webserver/error"
)

func Arg(index int, name ...string) (string, error.IError) {
	if len(os.Args) > index {
		return os.Args[index], error.NewError("")
	} else {
		var _name = "args[" + strconv.FormatInt(int64(index), 10) + "]"
		if len(name) > 0 {
			_name = name[0]
		}

		return "", error.NewError(_name + " not found !")
	}
}

func CharBack() string {
	return "\n"
}

func Tab() string {
	return "\t"
}

func Slash() string {
	var slash = "/"

	if runtime.GOOS == "windows" {
		slash = "\\"
	}

	return slash
}

func PwdVar() string {
	start := "("
	end := ")"

	if runtime.GOOS == "windows" {
		start = "{"
		end = "}"
	}

	return "$" + start + "pwd" + end
}

func RootPath() string {
	var root = "/"

	if runtime.GOOS == "windows" {
		root = "C:\\"
	}

	return root
}

func IsBuild() bool {
	return os.Getenv("GOBUILD") != "1"
}

type String struct {
	String string
}

func (s *String) IsError() bool {
	return s.String != "" || s.String == "ERROR : "
}

func (s *String) Append(str string) *String {
	(*s).String += str

	return s
}

func (s *String) AppendIf(condition bool, ifTrue string, ifFalse string) *String {
	if condition {
		(*s).String += ifTrue
	} else {
		(*s).String += ifFalse
	}

	return s
}

func (s *String) IsEmpty() bool {
	return strings.Trim(s.String, " ") == ""
}

func RandomNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}

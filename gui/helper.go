package main

import (
	"fmt"
	"github.com/asticode/go-astilectron"
	"io"
	"log"
	"npbg/gui/writer"
	"npbg/helpers"
	"npbg/http/portChoice"
	"os"
	"strconv"
	"strings"
)

func UrlBase() string {
	port := strconv.FormatInt(int64(portChoice.ChosenPort), 10)

	return "http://127.0.0.1:" + port
}

func IsDevEnv() bool {
	return os.Getenv("OPEN_DEVTOOLS") == "1" ||
		strings.Contains(os.Args[0], helpers.Slash()+"b001"+helpers.Slash()+"exe"+helpers.Slash())
}

func OpenDevTools(w *astilectron.Window, l *log.Logger) {
	if IsDevEnv() {
		err := w.OpenDevTools()
		if err != nil {
			l.Fatal(fmt.Errorf("erreur : %s", err.Error()))
		}
	}
}

func GetLoggerWriter() io.Writer {
	if IsDevEnv() {
		return writer.NewLogWriter().Enable().Writer
	}
	return writer.NewLogWriter().Disable().Writer
}

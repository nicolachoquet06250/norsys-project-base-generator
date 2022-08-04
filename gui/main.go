package main

import (
	_ "embed"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"npbg/helpers"
	"npbg/history"
	"npbg/http/portChoice"
	"npbg/server"
	"os"
	"strconv"
	"strings"
)

//go:embed logo-norsys.png
var icon string

var firstLoad = true

func main() {
	var openDevTools = os.Getenv("OPEN_DEVTOOLS") == "1" || os.Getenv("OPEN_DEVTOOLS") == "" ||
		!strings.Contains(os.Args[0], helpers.Slash()+"b001"+helpers.Slash()+"exe"+helpers.Slash())

	go server.Process(true, false)

	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	GenerateIcon()

	a := CreateApp(l, "NPBG")
	defer a.Close()

	Loader := CreateLoader(a, l)

	urlBase := "http://127.0.0.1:" + strconv.FormatInt(int64(portChoice.ChosenPort), 10)

	w := CreateWindow(
		a, l,
		urlBase,
		&astilectron.WindowOptions{
			Center: astikit.BoolPtr(true),
			Height: astikit.IntPtr(700),
			Width:  astikit.IntPtr(700),
			Icon:   astikit.StrPtr(history.GetIconPath()),
			Show:   astikit.BoolPtr(false),
		},
	)

	if openDevTools {
		err := w.OpenDevTools()
		if err != nil {
			println(err.Error())
		}
	}

	w.On(astilectron.EventNameWindowEventReadyToShow, func(e astilectron.Event) (deleteListener bool) {
		if firstLoad == true && Loader != nil {
			err := Loader.Hide()
			println("destroy loader window")
			if err != nil {
				l.Fatal(fmt.Errorf("loader window can't be destroy"))
			}
		}

		err := w.Show()
		println("show main window")
		if err != nil {
			l.Fatal(fmt.Errorf("main window can't be showed"))
		}

		firstLoad = false
		return
	})

	w.On(astilectron.EventNameWindowEventResize, func(e astilectron.Event) (deleteListener bool) {
		println("Window resized")
		return
	})

	w.On(astilectron.EventNameWindowEventClosed, func(e astilectron.Event) (deleteListener bool) {
		_ = Loader.Destroy()
		return
	})

	w.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		var jsonMessage JsonMessage = decodeJsonMessage(
			decodeMessage(m),
		)

		if jsonMessage.Channel == "Notification" {
			receiveNotificationChannel(a, w, &jsonMessage)
		}

		return
	})

	// Blocking pattern
	a.Wait()
}

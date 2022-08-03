package main

import (
	_ "embed"
	"encoding/base64"
	jsonPkg "encoding/json"
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

	err := w.OpenDevTools()
	if err != nil {
		println(err.Error())
	}

	if !openDevTools {
		err = w.CloseDevTools()
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
		json, err := m.MarshalJSON()
		if err != nil {
			log.Fatal(fmt.Printf("error : %s", err.Error()))
		}

		json, err = base64.StdEncoding.DecodeString(strings.Replace(string(json), "\"", "", 2))
		if err != nil {
			log.Fatal(fmt.Printf("error : %s", err.Error()))
		}
		strJson := strings.Replace(string(json), "\"{", "{", 1)
		strJson = strings.Replace(strJson, "}\"", "}", 1)
		strJson = strings.Replace(strJson, "\\\"", "\"", -1)

		println("message received decoded : " + strJson)

		var jsonMessage JsonMessage
		err = jsonPkg.Unmarshal([]byte(strJson), &jsonMessage)
		if err != nil {
			log.Fatal(fmt.Printf("error : %s", err.Error()))
		}

		if jsonMessage.Channel == "Notification" {
			notification := CreateNotification(a, NotificationOption{
				Title:    "Test notif",
				Subtitle: astikit.StrPtr("Test notif"),
				Body:     fmt.Sprintf("Bonjour\n%s", jsonMessage.Data["name"]),
			})

			_ = notification.Show()

			notification.On(astilectron.EventNameNotificationEventClicked, func(e astilectron.Event) (deleteListener bool) {
				err = w.SendMessage(JsonMessage{
					Channel: "Redirect",
					Data:    map[string]string{"uri": "/help"},
				})
				if err != nil {
					log.Fatal(fmt.Printf("error : %s", err.Error()))
				}
				return
			})
		}

		return
	})

	// Blocking pattern
	a.Wait()
}

package main

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"npbg/files"
	"npbg/helpers"
	"npbg/history"
	"npbg/http/portChoice"
	"npbg/http/routing"
	"npbg/notifications"
	"strconv"
	"strings"
)

var GeneratedPreventionAlert *helpers.Alert
var GeneratedAlert *helpers.Alert

func init() {
	GeneratedPreventionAlert = new(helpers.Alert)
	GeneratedAlert = new(helpers.Alert)
}

func GenerateIcon() {
	_icon := files.NewFile(history.GetIconPath())

	exists, _ := _icon.Exists()
	if !exists {
		err := _icon.Create(icon, true)
		if err != nil {
			log.Fatal(fmt.Errorf("can't create icon locally"))
		}
	}
}

func CreateApp(name string, l *log.Logger) (a *astilectron.Astilectron) {
	// Initialize astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           name,
		BaseDirectoryPath: "gui",
	})

	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}

	// Add a listener on Astilectron
	a.On(astilectron.EventNameAppCrash, func(e astilectron.Event) (deleteListener bool) {
		l.Println("App has crashed")
		return
	})

	// Handle signals
	a.HandleSignals()

	// Start
	err = a.Start()
	if err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	return
}

func CreateWindow(a *astilectron.Astilectron, l *log.Logger, url string, options *astilectron.WindowOptions, name ...string) (w *astilectron.Window) {
	if len(name) == 0 {
		name = append(name, "main")
	}

	w, err := a.NewWindow(url, options)
	if err != nil {
		l.Fatal(fmt.Errorf("%s: new window failed: %w", name[0], err))
	}

	err = w.Create()
	if err != nil {
		l.Fatal(fmt.Errorf("%s: creating window failed: %w", name[0], err))
	}

	return
}

func CreateLoader(a *astilectron.Astilectron, l *log.Logger) (w *astilectron.Window) {
	l.Println("------------------------ LOADER ------------------------")
	baseUrl := "http://127.0.0.1:" + strconv.FormatInt(int64(portChoice.ChosenPort), 10)
	w = CreateWindow(
		a, l,
		baseUrl+routing.RouteToString(routing.LoaderPage),
		&astilectron.WindowOptions{
			Frame:          astikit.BoolPtr(false),
			Center:         astikit.BoolPtr(true),
			Width:          astikit.IntPtr(200),
			Height:         astikit.IntPtr(200),
			Resizable:      astikit.BoolPtr(false),
			Fullscreenable: astikit.BoolPtr(false),
			Icon:           astikit.StrPtr(history.GetIconPath()),
			Transparent:    astikit.BoolPtr(true),
		},
		"loader",
	)

	return
}

func CreateNotification(l *log.Logger, o NotificationOption) {
	err := notifications.Notify(
		o.Title, o.Body,
		history.GetIconPath(),
		astikit.StrPtr("Norsys Project Base Generator"),
	)
	if err != nil {
		if strings.Contains(err.Error(), "not compatible") {
			GeneratedPreventionAlert = GenerateAlert(err.Error(), helpers.ERROR)

			message := o.Title + "<br />" + o.Body
			var status helpers.AlertStatus
			if strings.Contains(strings.ToLower(message), "succès") {
				status = helpers.SUCCESS
			} else if strings.Contains(strings.ToLower(message), "erreur") {
				status = helpers.ERROR
			} else {
				status = helpers.INFO
			}

			GeneratedAlert = GenerateAlert(message, status)
		} else {
			l.Fatal(fmt.Errorf("error : %s", err))
		}
	}
}

func GenerateAlert(message string, status helpers.AlertStatus) *helpers.Alert {
	a := helpers.NewAlert(message, status)
	return &a
}

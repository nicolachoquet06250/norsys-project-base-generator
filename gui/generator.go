package main

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"npbg/files"
	"npbg/history"
	"npbg/http/portChoice"
	"npbg/http/routing"
	"strconv"
)

func toAstilectronNotificationOptions(o *NotificationOption) (options *astilectron.NotificationOptions) {
	options = &astilectron.NotificationOptions{
		Title:            o.Title,
		Icon:             history.GetIconPath(),
		Body:             o.Body,
		ReplyPlaceholder: "type your reply here", // Only MacOSX
		HasReply:         astikit.BoolPtr(true),  // Only MacOSX
	}
	if o.Subtitle != nil {
		options = &astilectron.NotificationOptions{
			Title:            o.Title,
			Subtitle:         *o.Subtitle,
			Icon:             history.GetIconPath(),
			Body:             o.Body,
			ReplyPlaceholder: "type your reply here", // Only MacOSX
			HasReply:         astikit.BoolPtr(true),  // Only MacOSX
		}
	}

	return options
}

func GenerateIcon() {
	_icon := files.NewFile(history.GetIconPath())

	exists, _ := _icon.Exists()
	if !exists {
		err := _icon.Create(icon, true)
		if err != nil {
			log.Fatal(fmt.Printf("can't create icon locally"))
		}
	}
}

func CreateApp(l *log.Logger, name string) *astilectron.Astilectron {
	// Initialize astilectron
	var a, err = astilectron.New(l, astilectron.Options{
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

	return a
}

func CreateWindow(a *astilectron.Astilectron, l *log.Logger, url string, options *astilectron.WindowOptions, name ...string) *astilectron.Window {
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

	return w
}

func CreateLoader(a *astilectron.Astilectron, l *log.Logger) *astilectron.Window {
	l.Println("------------------------ LOADER ------------------------")
	baseUrl := "http://127.0.0.1:" + strconv.FormatInt(int64(portChoice.ChosenPort), 10)
	var w = CreateWindow(
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

	return w
}

func CreateNotification(a *astilectron.Astilectron, l *log.Logger, o NotificationOption) *astilectron.Notification {
	n := a.NewNotification(
		toAstilectronNotificationOptions(&o),
	)

	err := n.Create()

	if err != nil {
		l.Fatal(fmt.Errorf("erreur: new notification failed: %w", err))
	}

	return n
}

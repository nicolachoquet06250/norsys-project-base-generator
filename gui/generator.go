package main

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"npbg/files"
	"npbg/history"
	"npbg/http/portChoice"
	"strconv"
)

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
		println("App has crashed")
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
	baseUrl := "http://127.0.0.1:" + strconv.FormatInt(int64(portChoice.ChosenPort), 10)
	return CreateWindow(
		a, l,
		baseUrl+"/load",
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
}

func CreateNotification(a *astilectron.Astilectron, options NotificationOption) *astilectron.Notification {
	o := &astilectron.NotificationOptions{
		Title:            options.Title,
		Icon:             history.GetIconPath(),
		Body:             options.Body,
		ReplyPlaceholder: "type your reply here", // Only MacOSX
		HasReply:         astikit.BoolPtr(true),  // Only MacOSX
	}
	if options.Subtitle != nil {
		o = &astilectron.NotificationOptions{
			Title:            options.Title,
			Subtitle:         *options.Subtitle,
			Icon:             history.GetIconPath(),
			Body:             options.Body,
			ReplyPlaceholder: "type your reply here", // Only MacOSX
			HasReply:         astikit.BoolPtr(true),  // Only MacOSX
		}
	}

	n := a.NewNotification(o)

	// Add listeners
	n.On(astilectron.EventNameNotificationEventClicked, func(e astilectron.Event) (deleteListener bool) {
		log.Println("the notification has been clicked!")
		return
	})
	// Only for MacOSX
	n.On(astilectron.EventNameNotificationEventReplied, func(e astilectron.Event) (deleteListener bool) {
		log.Printf("the user has replied to the notification: %s\n", e.Reply)
		return
	})

	// Create notification
	_ = n.Create()

	return n
}

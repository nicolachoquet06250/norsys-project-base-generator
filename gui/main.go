package main

import (
	_ "embed"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"strconv"
	"test_go_webserver/http/portChoice"
	"test_go_webserver/server"
)

//go:embed logo-norsys.png
var icon string

var firstLoad = true
var Loader *astilectron.Window

func main() {
	go server.Process(true, false)

	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Initialize astilectron
	var a, err = astilectron.New(l, astilectron.Options{
		AppName:           "GUI test",
		BaseDirectoryPath: "gui",
	})

	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

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

	Loader, err := a.NewWindow("http://127.0.0.1:"+strconv.FormatInt(int64(portChoice.ChosenPort), 10)+"/load", &astilectron.WindowOptions{
		Frame:          astikit.BoolPtr(false),
		Center:         astikit.BoolPtr(true),
		Width:          astikit.IntPtr(200),
		Height:         astikit.IntPtr(200),
		Resizable:      astikit.BoolPtr(false),
		Fullscreenable: astikit.BoolPtr(false),
		Icon:           astikit.StrPtr(icon),
		Transparent:    astikit.BoolPtr(true),
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	err = Loader.Create()
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// Create a new window
	w, err := a.NewWindow("http://127.0.0.1:"+strconv.FormatInt(int64(portChoice.ChosenPort), 10), &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(700),
		Icon:   astikit.StrPtr(icon),
		Show:   astikit.BoolPtr(false),
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	err = w.Create()
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	w.On(astilectron.EventNameWindowEventReadyToShow, func(e astilectron.Event) (deleteListener bool) {
		if firstLoad == true && Loader != nil {
			err := Loader.Hide()
			println("destroy loader window")
			if err != nil {
				l.Fatal(fmt.Errorf("loader window can't be destroy"))
			}
		}

		err = w.Show()
		println("show main window")
		if err != nil {
			l.Fatal(fmt.Errorf("main window can't be showed"))
		}

		firstLoad = false
		return
	})

	// Add a listener on the window
	w.On(astilectron.EventNameWindowEventResize, func(e astilectron.Event) (deleteListener bool) {
		println("Window resized")
		return
	})

	w.On(astilectron.EventNameWindowEventClosed, func(e astilectron.Event) (deleteListener bool) {
		Loader.Destroy()
		return
	})

	// Blocking pattern
	a.Wait()
}

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

	// Handle signals
	a.HandleSignals()

	// Start
	err = a.Start()
	if err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// Create a new window
	w, err := a.NewWindow("http://127.0.0.1:"+strconv.FormatInt(int64(portChoice.ChosenPort), 10), &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(700),
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	err = w.Create()
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// Add a listener on Astilectron
	a.On(astilectron.EventNameAppCrash, func(e astilectron.Event) (deleteListener bool) {
		println("App has crashed")
		return
	})

	// Add a listener on the window
	w.On(astilectron.EventNameWindowEventResize, func(e astilectron.Event) (deleteListener bool) {
		println("Window resized")
		return
	})

	// Blocking pattern
	a.Wait()
}

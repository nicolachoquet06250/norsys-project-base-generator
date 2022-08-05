package main

import (
	_ "embed"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"npbg/history"
	"npbg/http/routing"
	"npbg/server"
)

//go:embed logo-norsys.png
var icon string

var Modal *astilectron.Window

func main() {
	go server.Process(true, false)

	logger := log.New(GetLoggerWriter(), log.Prefix(), log.Flags())

	GenerateIcon()

	app := CreateApp(logger, "NPBG")
	defer app.Close()

	urlBase := UrlBase()

	var Loader = CreateLoader(app, logger)

	var Window = CreateWindow(
		app, logger,
		urlBase,
		&astilectron.WindowOptions{
			Center: astikit.BoolPtr(true),
			Height: astikit.IntPtr(700),
			Width:  astikit.IntPtr(700),
			Icon:   astikit.StrPtr(history.GetIconPath()),
			Show:   astikit.BoolPtr(false),
		},
	)

	err := Window.Show()
	logger.Println("show main Window")
	if err != nil {
		logger.Fatal(fmt.Errorf("main Window can't be showed"))
	}

	_ = app.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astikit.StrPtr("Normal 1")},
				{Label: astikit.StrPtr("Normal 2")},
				{Type: astilectron.MenuItemTypeSeparator},
				{Label: astikit.StrPtr("Normal 3")},
			},
		},
		{
			Label: astikit.StrPtr("Aide"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("Aide"),
					Role:  astilectron.MenuItemRoleHelp,
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						var err = Window.SendMessage(
							NewMessage(Redirect, map[string]string{"uri": routing.RouteToString(routing.HelpPage)}),
						)
						if err != nil {
							log.Fatal(fmt.Printf("error : %s", err.Error()))
						}
						return
					},
				},
				{
					Label: astikit.StrPtr("Close"),
					Role:  astilectron.MenuItemRoleClose,
				},
			},
		},
	}).Create()

	OpenDevTools(Window, logger)

	Window.On(astilectron.EventNameWindowEventReadyToShow, func(e astilectron.Event) (deleteListener bool) {
		err = Window.Show()
		logger.Println("show main Window")
		if err != nil {
			logger.Fatal(fmt.Errorf("main Window can't be showed"))
		}
		return
	})

	Window.On(astilectron.EventNameWindowEventClosed, func(e astilectron.Event) (deleteListener bool) {
		if Loader != nil {
			_ = Loader.Destroy()
		}

		if Modal != nil {
			_ = Modal.Destroy()
		}
		return
	})

	Window.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		var jsonMessage = decodeJsonMessage(
			decodeMessage(m),
		)

		switch jsonMessage.Channel {
		case string(Notification):
			receiveNotificationChannel(app, Window, logger, &jsonMessage, nil)
			break
		case string(OpenFolderSelectorModal):
			Modal = receiveOpenFolderSelectorModalChannel(app, Window, logger, &jsonMessage, nil)
			break
		case string(DestroyLoader):
			receiveDestroyLoaderChannel(app, Loader, logger, &jsonMessage, Window)
			break
		}

		return
	})

	// Blocking pattern
	app.Wait()
}

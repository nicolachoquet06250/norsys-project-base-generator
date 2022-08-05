package main

import (
	_ "embed"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"io"
	"log"
	"npbg/gui/writer"
	"npbg/helpers"
	"npbg/history"
	"npbg/http/portChoice"
	"npbg/http/routing"
	"npbg/server"
	"os"
	"strconv"
	"strings"
)

//go:embed logo-norsys.png
var icon string

var firstLoad = true

func UrlBase() string {
	port := strconv.FormatInt(int64(portChoice.ChosenPort), 10)

	return "http://127.0.0.1:" + port
}

func OpenDevTools(w *astilectron.Window, l *log.Logger) {
	var isDevEnv = os.Getenv("OPEN_DEVTOOLS") == "1" ||
		!strings.Contains(os.Args[0], helpers.Slash()+"b001"+helpers.Slash()+"exe"+helpers.Slash())

	if isDevEnv {
		err := w.OpenDevTools()
		if err != nil {
			l.Fatal(fmt.Errorf("erreur : %s", err.Error()))
		}
	}
}

func main() {
	var isDevEnv = os.Getenv("OPEN_DEVTOOLS") == "1" ||
		!strings.Contains(os.Args[0], helpers.Slash()+"b001"+helpers.Slash()+"exe"+helpers.Slash())

	go server.Process(true, false)

	var _writer = (func() io.Writer {
		if isDevEnv {
			return writer.NewLogWriter().Enable().Writer
		}
		return writer.NewLogWriter().Disable().Writer
	})()

	logger := log.New(_writer, log.Prefix(), log.Flags())

	GenerateIcon()

	app := CreateApp(logger, "NPBG")
	defer app.Close()

	loader := CreateLoader(app, logger)

	urlBase := UrlBase()

	window := CreateWindow(
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

	var menu = app.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astikit.StrPtr("Normal 1")},
				{Label: astikit.StrPtr("Normal 2")},
				{Type: astilectron.MenuItemTypeSeparator},
				{Label: astikit.StrPtr("Normal 3")},
			},
		},
		/*{
			Label: astikit.StrPtr("Checkbox"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Checked: astikit.BoolPtr(true), Label: astikit.StrPtr("Checkbox 1"), Type: astilectron.MenuItemTypeCheckbox},
				{Label: astikit.StrPtr("Checkbox 2"), Type: astilectron.MenuItemTypeCheckbox},
				{Label: astikit.StrPtr("Checkbox 3"), Type: astilectron.MenuItemTypeCheckbox},
			},
		},
		{
			Label: astikit.StrPtr("Radio"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Checked: astikit.BoolPtr(true), Label: astikit.StrPtr("Radio 1"), Type: astilectron.MenuItemTypeRadio},
				{Label: astikit.StrPtr("Radio 2"), Type: astilectron.MenuItemTypeRadio},
				{Label: astikit.StrPtr("Radio 3"), Type: astilectron.MenuItemTypeRadio},
			},
		},*/
		{
			Label: astikit.StrPtr("Aide"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("Aide"),
					Role:  astilectron.MenuItemRoleHelp,
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						var err = window.SendMessage(
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
	})

	// Retrieve a menu item
	// This will retrieve the "Checkbox 1" item
	menuItem, _ := menu.Item(1, 0)

	// Add listener
	menuItem.On(astilectron.EventNameMenuItemEventClicked, func(e astilectron.Event) bool {
		logger.Printf("Menu item has been clicked. 'Checked' status is now %t", *e.MenuItemOptions.Checked)
		return false
	})

	// Create the menu
	_ = menu.Create()

	// Manipulate a menu item
	_ = menuItem.SetChecked(true)

	// Init a new menu item
	/*var newItem = menu.NewItem(&astilectron.MenuItemOptions{
		Label:   astikit.StrPtr("Inserted"),
		SubMenu: []*astilectron.MenuItemOptions{},
	})

	// Insert the menu item at position "1"
	_ = menu.Insert(1, newItem)*/

	// Fetch a sub menu
	subMenu, _ := menu.SubMenu(0)

	var Modal *astilectron.Window

	// Init a new menu item
	var newItem = subMenu.NewItem(&astilectron.MenuItemOptions{
		Label: astikit.StrPtr("Appended"),
		SubMenu: []*astilectron.MenuItemOptions{
			{Label: astikit.StrPtr("Appended 1")},
			{Label: astikit.StrPtr("Appended 2"), OnClick: func(e astilectron.Event) (deleteListener bool) {
				Modal = CreateWindow(app, logger, urlBase+routing.RouteToString(routing.FolderSelectorPage), &astilectron.WindowOptions{
					Center: astikit.BoolPtr(true),
					Height: astikit.IntPtr(700),
					Width:  astikit.IntPtr(700),
					Icon:   astikit.StrPtr(history.GetIconPath()),
					Show:   astikit.BoolPtr(false),
					Modal:  astikit.BoolPtr(true),
				})

				err := Modal.OpenDevTools()
				if err != nil {
					logger.Fatal(fmt.Errorf("erreur : %s", err.Error()))
				}

				_ = Modal.Show()

				Modal.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
					var jsonMessage = decodeJsonMessage(
						decodeMessage(m),
					)

					switch jsonMessage.Channel {
					case string(ChooseFolder):
						receiveChooseFolderChannel(app, Modal, logger, &jsonMessage, window)
						break
					}

					return
				})
				return
			}},
		},
	})

	// Append menu item dynamically
	_ = subMenu.Append(newItem)

	OpenDevTools(window, logger)

	window.On(astilectron.EventNameWindowEventReadyToShow, func(e astilectron.Event) (deleteListener bool) {
		if firstLoad == true && loader != nil {
			err := loader.Hide()
			logger.Println("destroy loader window")
			if err != nil {
				logger.Fatal(fmt.Errorf("loader window can't be destroy"))
			}
		}

		err := window.Show()
		logger.Println("show main window")
		if err != nil {
			logger.Fatal(fmt.Errorf("main window can't be showed"))
		}

		firstLoad = false
		return
	})

	window.On(astilectron.EventNameWindowEventClosed, func(e astilectron.Event) (deleteListener bool) {
		_ = loader.Destroy()

		if Modal != nil {
			_ = Modal.Destroy()
		}
		return
	})

	window.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		var jsonMessage = decodeJsonMessage(
			decodeMessage(m),
		)

		switch jsonMessage.Channel {
		case string(Notification):
			receiveNotificationChannel(app, window, logger, &jsonMessage, nil)
			break
		case string(OpenFolderSelectorModal):
			Modal = receiveOpenFolderSelectorModalChannel(app, window, logger, &jsonMessage, nil)
			break
		}

		return
	})

	// Blocking pattern
	app.Wait()
}

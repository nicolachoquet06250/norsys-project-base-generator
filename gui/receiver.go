package main

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"io/ioutil"
	"log"
	"npbg/helpers"
	"npbg/history"
	"npbg/http/routing"
)

func receiveNotificationChannel(l *log.Logger, message *JsonMessage) {
	body := fmt.Sprintf("Bonjour\n%s", message.Data["name"])
	if message.Data["body"] != "" {
		body = message.Data["body"]
	}

	title := "Test notif"
	if message.Data["title"] != "" {
		title = message.Data["title"]
	}

	println("NOTIFICATION : " + title + ", " + body)

	CreateNotification(l, NotificationOption{
		Title: title,
		Body:  body,
	})
}

func receiveChooseFolderChannel(w *astilectron.Window, l *log.Logger, message *JsonMessage, main *astilectron.Window) {
	folderPath := message.Data["path"]

	CreateNotification(l, NotificationOption{
		Title: "Répertoire choisis",
		Body:  folderPath,
	})

	if w != nil {
		err := w.Destroy()
		if err != nil {
			l.Fatal(fmt.Errorf("fail to destroy window %s", err))
		}

		err = main.SendMessage(
			NewMessage(PutFolder, map[string]string{
				"folder": folderPath,
			}),
		)
		if err != nil {
			l.Fatal(fmt.Errorf("fail to send message to main window %s", err))
		}
	}
}

func receiveOpenFolderSelectorModalChannel(a *astilectron.Astilectron, w *astilectron.Window, l *log.Logger) *astilectron.Window {
	Modal = CreateWindow(a, l, UrlBase()+routing.RouteToString(routing.FolderSelectorPage), &astilectron.WindowOptions{
		Center:          astikit.BoolPtr(true),
		Height:          astikit.IntPtr(700),
		Width:           astikit.IntPtr(700),
		Icon:            astikit.StrPtr(history.GetIconPath()),
		Show:            astikit.BoolPtr(false),
		Modal:           astikit.BoolPtr(true),
		Transparent:     astikit.BoolPtr(false),
		BackgroundColor: astikit.StrPtr("white"),
	})

	OpenDevTools(Modal, l)

	err := Modal.Show()
	if err != nil {
		l.Fatal(fmt.Errorf("fail to show modal %s", err))
	}

	Modal.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		var jsonMessage = decodeJsonMessage(
			decodeMessage(m),
		)

		switch jsonMessage.Channel {
		case string(ChooseFolder):
			receiveChooseFolderChannel(Modal, l, &jsonMessage, w)
			break
		case string(OpenFolder):
			go receiveOpenFolderChannel(Modal, l, &jsonMessage)
			break
		}

		if GeneratedPreventionAlert != nil {
			err = w.SendMessage(NewMessage(ShowAlert, map[string]string{
				"message": GeneratedPreventionAlert.Message,
				"type":    string(GeneratedPreventionAlert.Type),
			}))
			if err != nil {
				log.Println(fmt.Errorf("main window send message error : %s", err))
			}
		}

		if GeneratedAlert != nil {
			err = w.SendMessage(NewMessage(ShowAlert, map[string]string{
				"message": GeneratedAlert.Message,
				"type":    string(GeneratedAlert.Type),
			}))
			if err != nil {
				log.Println(fmt.Errorf("main window send message error : %s", err))
			}
		}

		return
	})

	return Modal
}

func receiveOpenFolderChannel(w *astilectron.Window, l *log.Logger, message *JsonMessage) {
	if message.Data["folder"] == "" {
		message.Data["folder"] = helpers.HomePath()
	}

	files, err := ioutil.ReadDir(message.Data["folder"])
	if err != nil {
		log.Fatal(err)
	}

	var tree []string

	for _, f := range files {
		if f.IsDir() {
			tree = append(tree, f.Name())
		}
	}

	err = w.SendMessage(
		NewArrayMessage(GetTree, map[string]interface{}{
			"basePath":      message.Data["folder"],
			"tree":          tree,
			"pathSeparator": helpers.Slash(),
			"isHome":        message.Data["folder"] == helpers.HomePath(),
		}),
	)
	if err != nil {
		l.Fatal(fmt.Errorf("fail to send message to window %s", err))
	}
}

func receiveDestroyLoaderChannel(w *astilectron.Window, l *log.Logger) {
	if w != nil {
		err := w.Destroy()

		m := "destroy loader window"
		if err != nil {
			m = "loader window can't be destroy because doesn't exists"
		}

		l.Println(m)
		l.Println("------------------------ END LOADER ------------------------")
	}
}

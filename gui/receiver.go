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

func receiveNotificationChannel(a *astilectron.Astilectron, w *astilectron.Window, l *log.Logger, message *JsonMessage, main *astilectron.Window) {
	body := fmt.Sprintf("Bonjour\n%s", message.Data["name"])
	if message.Data["body"] != "" {
		body = message.Data["body"]
	}

	title := "Test notif"
	if message.Data["title"] != "" {
		title = message.Data["title"]
	}

	n := CreateNotification(a, l, NotificationOption{
		Title:    title,
		Subtitle: astikit.StrPtr(title),
		Body:     body,
	})

	_ = n.Show()
}

func receiveChooseFolderChannel(a *astilectron.Astilectron, w *astilectron.Window, l *log.Logger, message *JsonMessage, main *astilectron.Window) {
	folderPath := message.Data["path"]

	//println("path : " + folderPath)

	n := CreateNotification(a, l, NotificationOption{
		Title:    "Répertoire choisis",
		Subtitle: astikit.StrPtr("Répertoire choisis"),
		Body:     folderPath,
	})

	_ = n.Show()

	if w != nil {
		_ = w.Destroy()

		_ = main.SendMessage(
			NewMessage(PutFolder, map[string]string{
				"folder": folderPath,
			}),
		)
	}
}

func receiveOpenFolderSelectorModalChannel(a *astilectron.Astilectron, w *astilectron.Window, l *log.Logger, message *JsonMessage, main *astilectron.Window) *astilectron.Window {
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

	_ = Modal.Show()

	Modal.OnMessage(func(m *astilectron.EventMessage) (v interface{}) {
		var jsonMessage = decodeJsonMessage(
			decodeMessage(m),
		)

		switch jsonMessage.Channel {
		case string(ChooseFolder):
			receiveChooseFolderChannel(a, Modal, l, &jsonMessage, w)
			break
		case string(OpenFolder):
			go receiveOpenFolderChannel(a, Modal, l, &jsonMessage, w)
			break
		}

		return
	})

	return Modal
}

func receiveOpenFolderChannel(a *astilectron.Astilectron, w *astilectron.Window, l *log.Logger, message *JsonMessage, main *astilectron.Window) {
	if message.Data["folder"] == "" {
		message.Data["folder"] = helpers.HomePath()
	}

	println(message.Data["folder"])

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

	_ = w.SendMessage(
		NewArrayMessage(GetTree, map[string]interface{}{
			"tree": tree,
		}),
	)
}

func receiveDestroyLoaderChannel(a *astilectron.Astilectron, w *astilectron.Window, l *log.Logger, message *JsonMessage, main *astilectron.Window) {
	err := w.Destroy()
	l.Println("destroy Loader Window")
	l.Println("------------------------ END LOADER ------------------------")
	if err != nil {
		l.Fatal(fmt.Errorf("Loader Window can't be destroy"))
	}
}

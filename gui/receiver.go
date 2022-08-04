package main

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
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

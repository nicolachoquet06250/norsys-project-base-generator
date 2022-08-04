package main

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
)

func receiveNotificationChannel(a *astilectron.Astilectron, w *astilectron.Window, message *JsonMessage) {
	body := fmt.Sprintf("Bonjour\n%s", message.Data["name"])
	if message.Data["body"] != "" {
		body = message.Data["body"]
	}

	title := "Test notif"
	if message.Data["title"] != "" {
		title = message.Data["title"]
	}

	notification := CreateNotification(a, NotificationOption{
		Title:    title,
		Subtitle: astikit.StrPtr(title),
		Body:     body,
	})

	_ = notification.Show()

	notification.On(astilectron.EventNameNotificationEventClicked, func(e astilectron.Event) (deleteListener bool) {
		err := w.SendMessage(JsonMessage{
			Channel: "Redirect",
			Data:    map[string]string{"uri": "/help"},
		})
		if err != nil {
			log.Fatal(fmt.Printf("error : %s", err.Error()))
		}
		return
	})
}

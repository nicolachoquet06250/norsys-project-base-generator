package main

type Channel string

const (
	Redirect     Channel = "Redirect"
	Notification Channel = "Notification"
	ChooseFolder Channel = "ChooseFolder"
	PutFolder    Channel = "PutFolder"
)

type NotificationOption struct {
	Title    string
	Subtitle *string
	Body     string
}

type JsonMessage struct {
	Channel string            `json:"channel"`
	Data    map[string]string `json:"data"`
}

func NewMessage(channel Channel, data map[string]string) JsonMessage {
	return JsonMessage{
		Channel: string(channel),
		Data:    data,
	}
}

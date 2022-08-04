package main

type Channel string

const (
	Redirect                Channel = "Redirect"
	Notification            Channel = "Notification"
	ChooseFolder            Channel = "ChooseFolder"
	PutFolder               Channel = "PutFolder"
	OpenFolderSelectorModal Channel = "OpenFolderSelectorModal"
	OpenFolder              Channel = "OpenFolder"
	GetTree                 Channel = "GetTree"
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

type JsonArrayMessage struct {
	Channel string         `json:"channel"`
	Data    map[string]any `json:"data"`
}

func NewArrayMessage(channel Channel, data map[string]any) JsonArrayMessage {
	return JsonArrayMessage{
		Channel: string(channel),
		Data:    data,
	}
}

package main

type NotificationOption struct {
	Title    string
	Subtitle *string
	Body     string
}

type JsonMessage struct {
	Channel string            `json:"channel"`
	Data    map[string]string `json:"data"`
}

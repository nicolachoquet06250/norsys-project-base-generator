package main

import (
	"encoding/base64"
	jsonPkg "encoding/json"
	"fmt"
	"github.com/asticode/go-astilectron"
	"log"
	"strings"
)

func decodeMessage(m *astilectron.EventMessage) []byte {
	json, err := m.MarshalJSON()
	if err != nil {
		log.Fatal(fmt.Printf("error : %s", err.Error()))
	}

	json, err = base64.StdEncoding.DecodeString(strings.Replace(string(json), "\"", "", 2))
	if err != nil {
		log.Fatal(fmt.Printf("error : %s", err.Error()))
	}

	return json
}

func decodeJsonMessage(message []byte) (jsonMessage JsonMessage) {
	err := jsonPkg.Unmarshal(message, &jsonMessage)
	if err != nil {
		log.Fatal(fmt.Printf("error : %s", err.Error()))
	}

	return jsonMessage
}

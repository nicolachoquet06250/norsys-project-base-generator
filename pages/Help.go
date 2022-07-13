package pages

import (
	_ "embed"
	"net/http"
	"test_go_webserver/http/portChoice"
	. "test_go_webserver/pages/helpers"
)

//go:embed templates/help.html
var help string

func Help(w http.ResponseWriter, r *http.Request) {
	result, _ := ParsePage("help", help, map[string]interface{}{
		"Path": r.URL.Path,
		"Home": r.URL.Path == "/",
		"Port": portChoice.ChosenPort,
	})
	Text(&w, result)
}

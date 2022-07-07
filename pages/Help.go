package pages

import (
	_ "embed"
	"net/http"
)

//go:embed templates/help.html
var help string

func Help(w http.ResponseWriter, r *http.Request) {
	result, _ := ParsePage("help", help, map[string]interface{}{
		"Path": r.URL.Path,
		"Home": r.URL.Path == "/",
	})
	Text(&w, result)
}

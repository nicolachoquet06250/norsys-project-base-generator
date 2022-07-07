package pages

import (
	_ "embed"
	"net/http"
	"test_go_webserver/technos"
)

//go:embed templates/index.html
var home string

func Home(w http.ResponseWriter, r *http.Request) {
	result, _ := ParsePage("index", home, map[string]interface{}{
		"Path":    r.URL.Path,
		"Home":    r.URL.Path == "/",
		"Technos": technos.All(),
	})
	Text(&w, result)
}

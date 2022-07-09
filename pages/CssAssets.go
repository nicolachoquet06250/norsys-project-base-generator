package pages

import (
	_ "embed"
	"github.com/gorilla/mux"
	"net/http"
)

//go:embed templates/assets/home.css
var homeCss string

const (
	HOME = "home"
	VOID = ""
)

func CssAssets(w http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["file"]

	if fileName != VOID {
		file := func() string {
			switch fileName {
			case HOME:
				return homeCss
			default:
				return VOID
			}
		}()

		query := r.URL.Query()
		queryStringParams := map[string]interface{}{}

		for param := range query {
			queryStringParams[param] = query.Get(param)
		}

		result, _ := ParsePage("css", file, queryStringParams)

		Css(&w, result)
	} else {
		Css(&w, VOID)
	}
}

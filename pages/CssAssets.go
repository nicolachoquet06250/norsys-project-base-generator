package pages

import (
	_ "embed"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	. "test_go_webserver/pages/helpers"
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
				re := regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)\)`)
				substitution := "{{ .$variable }}"
				homeCss = re.ReplaceAllString(homeCss, substitution)

				re = regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)((, +)?(?P<defaultValue>[a-zA-Z-_'"]+))?\)`)
				substitution = "{{ if .$variable }} {{.$variable}} {{ else }}$defaultValue{{ end }}"
				homeCss = re.ReplaceAllString(homeCss, substitution)

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

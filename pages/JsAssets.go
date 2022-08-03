package pages

import (
	_ "embed"
	"github.com/gorilla/mux"
	"net/http"
	. "npbg/pages/helpers"
	"regexp"
)

//go:embed templates/assets/bootstrap/js/bootstrap.min.js
var bootstrapProdJs string

//go:embed templates/assets/bootstrap/js/bootstrap.js
var bootstrapHorsProdJs string

const (
	ASTILECTRON = "init_astilectron"
)

func BootstrapJsAssets(w http.ResponseWriter, r *http.Request) {
	if r.Host == "localhost" || r.Host == "127.0.0.1" {
		Js(&w, bootstrapHorsProdJs)
	} else {
		Js(&w, bootstrapProdJs)
	}
}

func JsAssets(w http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["file"]

	if fileName != VOID {
		file := func() string {
			switch fileName {
			case ASTILECTRON:
				re := regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)\)`)
				substitution := "{{ .$variable }}"
				initAstilectronJs = re.ReplaceAllString(initAstilectronJs, substitution)

				/*re = regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)((, +)?(?P<defaultValue>[a-zA-Z-_'"]+))?\)`)
				substitution = "{{ if .$variable }} {{.$variable}} {{ else }}$defaultValue{{ end }}"
				initAstilectronJs = re.ReplaceAllString(initAstilectronJs, substitution)*/

				return initAstilectronJs
			default:
				return VOID
			}
		}()

		query := r.URL.Query()
		queryStringParams := map[string]interface{}{}

		for param := range query {
			queryStringParams[param] = query.Get(param)
		}

		result, _ := ParsePage("js", file, &queryStringParams, "")

		Js(&w, result)
	} else {
		Js(&w, VOID)
	}
}

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

//go:embed templates/assets/init_astilectron.js
var initAstilectronJs string

//go:embed templates/assets/folder_selector.js
var folderSelectorJs string

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

				return initAstilectronJs
			case FOLDER_SELECROR:
				re := regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)\)`)
				substitution := "{{ .$variable }}"
				folderSelectorJs = re.ReplaceAllString(folderSelectorJs, substitution)

				return folderSelectorJs
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

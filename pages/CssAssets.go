package pages

import (
	_ "embed"
	"github.com/gorilla/mux"
	"net/http"
	. "npbg/pages/helpers"
	"regexp"
)

//go:embed templates/assets/home.css
var homeCss string

//go:embed templates/assets/generate.css
var generateCss string

//go:embed templates/assets/loader.css
var loaderCss string

//go:embed templates/assets/folder_selector.css
var folderSelectorCss string

//go:embed templates/assets/help.css
var helpCss string

//go:embed templates/assets/bootstrap/css/bootstrap.min.css
var bootstrapProdCss string

//go:embed templates/assets/bootstrap/css/bootstrap.css
var bootstrapHorsProdCss string

//go:embed templates/assets/bootstrap/css/bootstrap.css.map
var bootstrapHorsProdCssMap string

//go:embed templates/assets/bootstrap/css/bootstrap.min.css.map
var bootstrapProdCssMap string

//go:embed templates/assets/custom-highlightjs-line-numbers.css
var customHighlightjsCssMap string

const (
	HOME               = "home"
	LOADER             = "loader"
	GENERATE           = "generate"
	HELP               = "help"
	FOLDER_SELECROR    = "folder_selector"
	CUSTOM_HIGHLIGHTJS = "custom-highlightjs-line-numbers"
	VOID               = ""
)

func getCssContentFromSource(source string) string {
	re := regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)\)`)
	substitution := "{{ .$variable }}"
	source = re.ReplaceAllString(source, substitution)

	re = regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)((, +)?(?P<defaultValue>[a-zA-Z-_'"]+))?\)`)
	substitution = "{{ if .$variable }} {{.$variable}} {{ else }}$defaultValue{{ end }}"
	source = re.ReplaceAllString(source, substitution)

	return source
}

func CssAssets(w http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["file"]

	if fileName != VOID {
		file := func() string {
			switch fileName {
			case HOME:
				return getCssContentFromSource(homeCss)
			case LOADER:
				return getCssContentFromSource(loaderCss)
			case GENERATE:
				return getCssContentFromSource(generateCss)
			case FOLDER_SELECROR:
				return getCssContentFromSource(folderSelectorCss)
			case CUSTOM_HIGHLIGHTJS:
				return getCssContentFromSource(customHighlightjsCssMap)
			case HELP:
				return getCssContentFromSource(helpCss)
			default:
				return VOID
			}
		}()

		query := r.URL.Query()
		queryStringParams := map[string]interface{}{}

		for param := range query {
			queryStringParams[param] = query.Get(param)
		}

		result, _ := ParsePage("css", file, &queryStringParams, "")

		Css(&w, result)
	} else {
		Css(&w, VOID)
	}
}

func BootstrapCssAssets(w http.ResponseWriter, r *http.Request) {
	if r.Host == "localhost" || r.Host == "127.0.0.1" {
		Css(&w, bootstrapHorsProdCss)
	} else {
		Css(&w, bootstrapProdCss)
	}
}

func BootstrapCssMapAssets(w http.ResponseWriter, r *http.Request) {
	if r.Host == "localhost" || r.Host == "127.0.0.1" {
		Css(&w, bootstrapHorsProdCssMap)
	} else {
		Css(&w, bootstrapProdCssMap)
	}
}

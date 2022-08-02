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

//go:embed templates/assets/generate.css
var generateCss string

//go:embed templates/assets/loader.css
var loaderCss string

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

const (
	HOME     = "home"
	LOADER   = "loader"
	GENERATE = "generate"
	HELP     = "help"
	VOID     = ""
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
			case LOADER:
				re := regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)\)`)
				substitution := "{{ .$variable }}"
				loaderCss = re.ReplaceAllString(loaderCss, substitution)

				re = regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)((, +)?(?P<defaultValue>[a-zA-Z-_'"]+))?\)`)
				substitution = "{{ if .$variable }} {{.$variable}} {{ else }}$defaultValue{{ end }}"
				loaderCss = re.ReplaceAllString(loaderCss, substitution)

				return loaderCss
			case GENERATE:
				re := regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)\)`)
				substitution := "{{ .$variable }}"
				generateCss = re.ReplaceAllString(generateCss, substitution)

				re = regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)((, +)?(?P<defaultValue>[a-zA-Z-_'"]+))?\)`)
				substitution = "{{ if .$variable }} {{.$variable}} {{ else }}$defaultValue{{ end }}"
				generateCss = re.ReplaceAllString(generateCss, substitution)

				return generateCss
			case HELP:
				re := regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)\)`)
				substitution := "{{ .$variable }}"
				helpCss = re.ReplaceAllString(helpCss, substitution)

				re = regexp.MustCompile(`(?m)go-bind\((?P<variable>[a-zA-Z-_]+)((, +)?(?P<defaultValue>[a-zA-Z-_'"]+))?\)`)
				substitution = "{{ if .$variable }} {{.$variable}} {{ else }}$defaultValue{{ end }}"
				helpCss = re.ReplaceAllString(helpCss, substitution)

				return helpCss
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

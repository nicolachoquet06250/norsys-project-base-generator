package pages

import (
	_ "embed"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "npbg/configFiles"
	. "npbg/pages/helpers"
	"npbg/technos"
)

//go:embed templates/index.html
var home string

//go:embed templates/menu.html
var Menu string

//go:embed templates/loader.html
var load string

//go:embed templates/folder_selector.html
var FolderSelector string

func Home(w http.ResponseWriter, r *http.Request) {
	result, err := ShowHtmlPage(Page{
		CurrentPage: "index",
		Template:    home,
		Title: Title{
			Tab:  "Formulaire de génération d'environements de dev",
			Page: "Norsys Project Base Generator",
		},
		CssFiles: CssFiles{
			"assets/home.css",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css",
			"https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.6.0/styles/default.min.css",
			"assets/custom-highlightjs-line-numbers.css",
		},
		MetaData: MetaData{
			Meta{
				Charset:   "UTF-8",
				Name:      "",
				Content:   "",
				HttpEquiv: "",
			},
			Meta{
				Charset:   "",
				Name:      "viewport",
				Content:   "width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0",
				HttpEquiv: "",
			},
			Meta{
				Charset:   "",
				Name:      "",
				Content:   "ie=edge",
				HttpEquiv: "X-UA-Compatible",
			},
		},
		Vars: &map[string]interface{}{
			"ProjectPath": nil,
			"ProjectName": nil,
			"Technos":     GetTechnoList(nil),
			"IsGenerate":  false,
			"Alert":       nil,
		},
	}, Menu)

	if err != nil {
		log.Fatal(err.Error())
	}

	Text(&w, result)
}

func Loader(w http.ResponseWriter, r *http.Request) {
	result, _ := ShowHtmlPage(Page{
		CurrentPage: "load",
		Template:    load,
		CssFiles:    CssFiles{"assets/loader.css"},
		MetaData: MetaData{
			Meta{
				Charset:   "UTF-8",
				Name:      "",
				Content:   "",
				HttpEquiv: "",
			},
			Meta{
				Charset:   "",
				Name:      "viewport",
				Content:   "width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0",
				HttpEquiv: "",
			},
			Meta{
				Charset:   "",
				Name:      "",
				Content:   "ie=edge",
				HttpEquiv: "X-UA-Compatible",
			},
		},
		Title: Title{
			Tab:  "Chargement...",
			Page: "Chargement...",
		},
		Vars: &map[string]interface{}{},
	}, Menu)

	Text(&w, result)
}

func GetFilesFromTechno(w http.ResponseWriter, r *http.Request) {
	techno, _ := technos.FromValue(mux.Vars(r)["techno"])

	j, _ := json.Marshal(ConfigFiles[techno.Name])

	Text(&w, string(j))
}

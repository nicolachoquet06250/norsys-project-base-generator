package pages

import (
	_ "embed"
	"log"
	"net/http"
	. "test_go_webserver/pages/helpers"
)

//go:embed templates/index.html
var home string

//go:embed templates/menu.html
var menu string

func Home(w http.ResponseWriter, r *http.Request) {
	result, err := ShowHtmlPage(Page{
		CurrentPage: "index",
		Template:    home,
		Title: Title{
			Tab:  "Formulaire de génération d'environements de dev",
			Page: "Norsys Project Base Generator",
		},
		CssFiles: CssFiles{"assets/home.css"},
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
	}, menu)

	if err != nil {
		log.Fatal(err.Error())
	}

	Text(&w, result)
}

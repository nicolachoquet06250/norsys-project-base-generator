package pages

import (
	_ "embed"
	"net/http"
	"test_go_webserver/http/portChoice"
	. "test_go_webserver/pages/helpers"
)

//go:embed templates/help.html
var help string

func Help(w http.ResponseWriter, r *http.Request) {
	result, _ := ShowHtmlPage(Page{
		CurrentPage: "help",
		Template:    help,
		Title: Title{
			Tab:  "Aide",
			Page: "Aide",
		},
		CssFiles: CssFiles{"assets/help.css"},
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
		Vars: &map[string]interface{}{"Port": portChoice.ChosenPort},
	}, menu)
	Text(&w, result)
}

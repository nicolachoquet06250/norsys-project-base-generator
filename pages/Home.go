package pages

import (
	_ "embed"
	"log"
	"net/http"
	. "test_go_webserver/pages/helpers"
)

//go:embed templates/index.html
var home string

func Home(w http.ResponseWriter, r *http.Request) {
	result, err := ParsePage("index", home, map[string]interface{}{
		"PageTitle":   nil,
		"CssFile":     "assets/home.css",
		"ProjectPath": nil,
		"ProjectName": nil,
		"Technos":     GetTechnoList(nil),
		"IsGenerate":  false,
		"Alert":       nil,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	Text(&w, result)
}

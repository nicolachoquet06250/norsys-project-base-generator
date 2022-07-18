package pages

import (
	_ "embed"
	"log"
	"net/http"
	"test_go_webserver/files"
	. "test_go_webserver/helpers"
	. "test_go_webserver/pages/helpers"
	"test_go_webserver/technos"
)

//go:embed templates/index.html
var generate string

func Generate(w http.ResponseWriter, r *http.Request) {
	projectName := r.URL.Query().Get("projectName")
	projectPath := r.URL.Query().Get("path") + Slash() + projectName
	technoValue := r.URL.Query().Get("techno")
	techno, technoErr := technos.FromValue(technoValue)

	if technoErr != nil {
		Text(&w, "Erreur "+technoErr.Error())
		return
	}

	alert := files.ProjectGeneration(projectPath, techno, &projectName)

	result, err := ParsePage("generate", generate, map[string]interface{}{
		"PageTitle":   "Génération du projet",
		"CssFile":     "assets/generate.css",
		"ProjectPath": projectPath,
		"ProjectName": projectName,
		"Technos":     GetTechnoList(&techno),
		"IsGenerate":  true,
		"Alert":       alert,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	Text(&w, result)
}

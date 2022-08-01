package pages

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"test_go_webserver/files"
	. "test_go_webserver/helpers"
	"test_go_webserver/history"
	. "test_go_webserver/pages/helpers"
	"test_go_webserver/technos"
)

//go:embed templates/index.html
var generate string

func Generate(w http.ResponseWriter, r *http.Request) {
	projectName := r.URL.Query().Get("projectName")
	projectPath := r.URL.Query().Get("path")
	technoValue := r.URL.Query().Get("techno")
	techno, technoErr := technos.FromValue(technoValue)

	if technoErr != nil {
		Text(&w, "Erreur "+technoErr.Error())
		return
	}

	project := files.NewProject(projectPath, &projectName, techno)

	var alert Alert
	exists, _ := project.Exists()
	if exists {
		alert = NewAlert(fmt.Sprintf("Le projet %s existe déjà dans le répertoire %s !", project.Name, project.Path), ERROR)
	} else {
		alert = project.Create()

		if alert.Type == SUCCESS {
			item := history.NewItem(project.Path, &project.Name, project.Techno)
			err := item.AddProject()
			if err != nil {
				alert = NewAlert(err.Error(), ERROR)
			}
		}
	}

	result, err := ShowHtmlPage(Page{
		CurrentPage: "generate",
		Template:    generate,
		Title: Title{
			Tab:  "Formulaire de génération d'environements de dev",
			Page: "Génération du projet",
		},
		CssFiles: CssFiles{"assets/generate.css"},
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
			"ProjectPath": projectPath,
			"ProjectName": projectName,
			"Technos":     GetTechnoList(&techno),
			"IsGenerate":  true,
			"Alert":       alert,
		},
	}, menu)

	if err != nil {
		log.Fatal(err.Error())
	}
	Text(&w, result)
}

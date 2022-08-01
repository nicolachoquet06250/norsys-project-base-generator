package pages

import (
	_ "embed"
	"encoding/json"
	"net/http"
	"test_go_webserver/files"
	"test_go_webserver/history"
	. "test_go_webserver/pages/helpers"
)

//go:embed templates/generated_projects.html
var generatedPage string

func MyGeneratedProjects(w http.ResponseWriter, r *http.Request) {
	h := history.GetHistoryList()

	result, _ := ShowHtmlPage(Page{
		CurrentPage: "generated",
		Template:    generatedPage,
		Title: Title{
			Tab:  "Mes projets générés",
			Page: "Projets générés",
		},
		CssFiles: CssFiles{
			"assets/help.css",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.1/css/all.min.css",
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
		Vars: &map[string]interface{}{"Projects": h},
	}, menu)

	Text(&w, result)
}

func RemoveHistoryProject(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var p history.ItemHistory
	err := decoder.Decode(&p)

	if err != nil {
		println(err.Error())
	}

	err = p.RemoveProject()

	if err != nil {
		println(err.Error())
	} else {
		files.NewProject(p.Path, &p.Name, p.Techno).Remove()
	}

	j, _ := json.Marshal(history.GetHistoryList())

	Text(&w, string(j))
}

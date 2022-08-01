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

	result, _ := ParsePage("generated", generatedPage, &map[string]interface{}{
		"Projects":      h,
		"EmptyProjects": h.IsEmpty(),
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

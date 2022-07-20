package pages

import (
	_ "embed"
	"net/http"
	"test_go_webserver/history"
	. "test_go_webserver/pages/helpers"
)

//go:embed templates/generated_projects.html
var generatedPage string

func MyGeneratedProjects(w http.ResponseWriter, r *http.Request) {
	result, _ := ParsePage("generated", generatedPage, map[string]interface{}{
		"Projects": history.GetHistoryList(),
	})

	Text(&w, result)
}

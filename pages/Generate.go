package pages

import (
	_ "embed"
	"net/http"
	"test_go_webserver/files"
	. "test_go_webserver/helpers"
	. "test_go_webserver/pages/helpers"
	"test_go_webserver/technos"
)

//go:embed templates/generate.html
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

	files.ProjectGeneration(projectPath, techno)

	result, _ := ParsePage("generate", generate, map[string]interface{}{
		"Path":        r.URL.Path,
		"Home":        r.URL.Path == "/",
		"ProjectPath": projectPath,
		"ProjectName": projectName,
		"Techno":      techno.Name,
	})
	Text(&w, result)
}

package pages

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func CssAssets(w http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["file"]

	if fileName != "" {
		file, fileErr := os.ReadFile("pages/templates/assets/" + fileName + ".css")

		if fileErr != nil {
			println(fileErr.Error())
		}

		query := r.URL.Query()
		queryStringParams := map[string]interface{}{}

		for param := range query {
			queryStringParams[param] = query.Get(param)
		}

		result, _ := ParsePage("css", string(file), queryStringParams)

		Css(&w, result)
	} else {
		Css(&w, "")
	}
}

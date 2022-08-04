package widgets

import (
	"net/http"
	"npbg/pages"
	"npbg/pages/helpers"
)

var folderSelector = pages.FolderSelector
var menu = pages.Menu

func FolderSelector(w http.ResponseWriter, r *http.Request) {
	result, _ := helpers.ShowHtmlPage(helpers.Page{
		CurrentPage: "folderSelector",
		Template:    folderSelector,
		Title: helpers.Title{
			Tab:  "Sélecteur de répertoire",
			Page: "Sélecteur de répertoire",
		},
		MetaData: helpers.MetaData{},
		CssFiles: helpers.CssFiles{},
		Vars:     helpers.VoidVars(),
	}, menu)

	helpers.Text(&w, result)
}

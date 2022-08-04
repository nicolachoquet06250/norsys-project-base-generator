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
			Tab: "Ouvrir",
		},
		MetaData: helpers.MetaData{},
		CssFiles: helpers.CssFiles{},
		Vars:     helpers.VoidVars(),
	}, menu)

	helpers.Text(&w, result)
}

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
		MetaData: helpers.MetaData{
			helpers.Meta{
				Charset:   "UTF-8",
				Name:      "",
				Content:   "",
				HttpEquiv: "",
			},
			helpers.Meta{
				Charset:   "",
				Name:      "viewport",
				Content:   "width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0",
				HttpEquiv: "",
			},
			helpers.Meta{
				Charset:   "",
				Name:      "",
				Content:   "ie=edge",
				HttpEquiv: "X-UA-Compatible",
			},
		},
		CssFiles: helpers.CssFiles{
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.1.2/css/all.min.css",
			"assets/folder_selector.css",
		},
		Vars: helpers.VoidVars(),
	}, menu)

	helpers.Text(&w, result)
}

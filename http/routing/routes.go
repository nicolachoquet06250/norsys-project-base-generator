package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"npbg/http/httpMethods"
	. "npbg/pages"
	"npbg/pages/widgets"
)

type Route string

const (
	HomePage               Route = "/"
	LoaderPage             Route = "/load"
	GeneratePage           Route = "/generate"
	HelpPage               Route = "/help"
	GeneratedPage          Route = "/generated"
	FolderSelectorPage     Route = "/folderSelector"
	GetFilesFromTechnoPage Route = "/getFilesFromTechno/{techno:[a-zA-Z_-]+}"
)

func RouteToString(route Route) string {
	return string(route)
}

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc(RouteToString(HomePage), Home).
		Methods(httpMethods.GET)
	r.HandleFunc(RouteToString(LoaderPage), Loader).
		Methods(httpMethods.GET)
	r.HandleFunc(RouteToString(GeneratePage), Generate).
		Methods(httpMethods.GET)
	r.HandleFunc(RouteToString(HelpPage), Help).
		Methods(httpMethods.GET)
	r.HandleFunc(RouteToString(GeneratedPage), MyGeneratedProjects).
		Methods(httpMethods.GET)
	r.HandleFunc(RouteToString(GeneratedPage), RemoveHistoryProject).
		Methods(httpMethods.DELETE)

	r.HandleFunc(RouteToString(GetFilesFromTechnoPage), GetFilesFromTechno).
		Methods(httpMethods.GET)

	r.HandleFunc("/assets/{file:[a-z_-]+}.css", CssAssets).
		Methods(httpMethods.GET)
	r.HandleFunc("/assets/{file:[a-z_-]+}.js", JsAssets).
		Methods(httpMethods.GET)
	r.HandleFunc("/bootstrap/b.css", BootstrapCssAssets).
		Methods(httpMethods.GET)
	r.HandleFunc("/bootstrap/bootstrap.min.css.map", BootstrapCssMapAssets).
		Methods(httpMethods.GET)
	r.HandleFunc("/bootstrap/b.js", BootstrapJsAssets).
		Methods(httpMethods.GET)

	// WIDGETS
	r.HandleFunc(RouteToString(FolderSelectorPage), widgets.FolderSelector).
		Methods(httpMethods.GET)

	http.Handle("/", r)
}

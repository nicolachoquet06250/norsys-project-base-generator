package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"npbg/http/httpMethods"
	. "npbg/pages"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home).
		Methods(httpMethods.GET)
	r.HandleFunc("/load", Loader).
		Methods(httpMethods.GET)
	r.HandleFunc("/generate", Generate).
		Methods(httpMethods.GET)
	r.HandleFunc("/help", Help).
		Methods(httpMethods.GET)
	r.HandleFunc("/generated", MyGeneratedProjects).
		Methods(httpMethods.GET)
	r.HandleFunc("/generated", RemoveHistoryProject).
		Methods(httpMethods.DELETE)
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

	http.Handle("/", r)
}

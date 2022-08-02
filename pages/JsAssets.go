package pages

import (
	_ "embed"
	"net/http"
	. "npbg/pages/helpers"
)

//go:embed templates/assets/bootstrap/js/bootstrap.min.js
var bootstrapProdJs string

//go:embed templates/assets/bootstrap/js/bootstrap.js
var bootstrapHorsProdJs string

func BootstrapJsAssets(w http.ResponseWriter, r *http.Request) {
	if r.Host == "localhost" || r.Host == "127.0.0.1" {
		Js(&w, bootstrapHorsProdJs)
	} else {
		Js(&w, bootstrapProdJs)
	}
}

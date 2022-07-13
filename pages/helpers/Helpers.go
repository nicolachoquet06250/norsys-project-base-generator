package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func ParsePage(name string, tpl string, vars map[string]interface{}) (string, error) {
	t, err := template.New(name + ".html").Parse(tpl)
	if err != nil {
		return "", err
	}

	tmpWriter := new(strings.Builder)
	err = t.Execute(tmpWriter, vars)
	if err != nil {
		return "", err
	}

	return tmpWriter.String(), nil
}

func ParsePageWOVars(name string, tpl string) (string, error) {
	return ParsePage(name, tpl, map[string]interface{}{})
}

func Text(w *http.ResponseWriter, r string) {
	_, err := (*w).Write([]byte(r))
	if err != nil {
		_ = fmt.Errorf("error : %s", err)
	}
}

func Css(w *http.ResponseWriter, r string) {
	(*w).Header().Set("Content-Type", "text/css")

	_, err := (*w).Write([]byte(r))
	if err != nil {
		_ = fmt.Errorf("error : %s", err)
	}
}

func Js(w *http.ResponseWriter, r string) {
	(*w).Header().Set("Content-Type", "application/javascript")

	_, err := (*w).Write([]byte(r))
	if err != nil {
		_ = fmt.Errorf("error : %s", err)
	}
}

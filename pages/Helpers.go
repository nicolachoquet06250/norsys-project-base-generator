package pages

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

func Text(w *http.ResponseWriter, r string) {
	_, err := (*w).Write([]byte(r))
	if err != nil {
		_ = fmt.Errorf("error : %s", err)
	}
}

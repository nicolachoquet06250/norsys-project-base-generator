package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"test_go_webserver/technos"
)

func ParseString(name string, str string, vars map[string]interface{}) (r string, e error) {
	t, err := template.New(name).Parse(str)
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

type PotentiallySelectedTechno struct {
	technos.Techno
	Selected bool
}

func GetTechnoList(techno *technos.Techno) []PotentiallySelectedTechno {
	list := technos.All()
	var finalList []PotentiallySelectedTechno

	for _, v := range list {
		t := PotentiallySelectedTechno{
			Techno: v,
			Selected: func() bool {
				if techno == nil {
					return false
				}

				if (*techno).Value == v.Value {
					return true
				} else {
					return false
				}
			}(),
		}

		finalList = append(finalList, t)
	}

	return finalList
}

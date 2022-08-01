package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"test_go_webserver/helpers"
	"test_go_webserver/technos"
)

func ParseString(name string, str string, vars *map[string]interface{}) (r string, e error) {
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

func ParseMenu(currentPage string, menu string) (template.HTML, error) {
	m, err := template.New("menu.html").Parse(menu)
	if err != nil {
		return "", err
	}

	tmpWriter := new(strings.Builder)
	err = m.Execute(tmpWriter, map[string]interface{}{
		"CurrentPage": currentPage,
	})
	if err != nil {
		return "", err
	}

	return template.HTML(tmpWriter.String()), nil
}

func ParsePage(name string, tpl string, vars *map[string]interface{}, menu string) (string, error) {
	(*vars)["Menu"], _ = ParseMenu(name, menu)

	t, err := template.New(name + ".html").Parse(tpl)
	if err != nil {
		return "", err
	}

	tmpWriter := new(strings.Builder)
	err = t.Execute(tmpWriter, &vars)
	if err != nil {
		return "", err
	}

	return tmpWriter.String(), nil
}

func ParsePageWOVars(name string, tpl string, menu string) (string, error) {
	return ParsePage(name, tpl, &map[string]interface{}{}, menu)
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
	list := technos.AllAvailable()
	list = helpers.ArrayReverse(list)
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

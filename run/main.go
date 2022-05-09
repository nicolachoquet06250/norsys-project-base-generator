package run

import (
	"github.com/gookit/color"
	. "go.pkg/nchoquet/autocomplete"
	"go.pkg/nchoquet/error"
	. "go.pkg/nchoquet/helpers"
	"os"
	"strings"
)

func OnRun(cb func(techno string, projectPath string)) {
	var techno, projectPath string
	var technoErr, projectPathErr error.IError

	techno, technoErr = Arg(1, "techno")
	projectPath, projectPathErr = Arg(2, "projectPath")

	var technoQuestion = "Quelle techno utilise-tu ?"
	var technoReader = CreateAutoCompleteLineReader(TechnoCompleter)
	defer technoReader.Close()

	var projectPathQuestion = "Quel est le chemin de ton projet ?"
	var projectPathReader = CreateAutoCompleteLineReader(PathCompleter)
	defer projectPathReader.Close()

	var err = ""

	if technoErr.IsError() || projectPathErr.IsError() {
		if technoErr.IsError() {
			println(technoQuestion)
			line, rlErr := technoReader.Readline()

			if rlErr != nil {
				if err != "" {
					err += CharBack() + Tab()
				} else {
					err += "ERROR : "
				}
				err += technoErr.Error()
			} else {
				if strings.TrimSpace(line) == "" {
					if err != "" {
						err += CharBack() + Tab()
					} else {
						err += "ERROR : "
					}
					err += technoErr.Error()
				} else {
					var pwd, _ = os.Getwd()
					techno = line
					for _, el := range []string{"$(pwd)", "${pwd}"} {
						techno = strings.Replace(techno, el, pwd, -1)
					}
					techno = strings.Replace(techno, " "+Slash()+" ", Slash(), -1)
					techno = strings.TrimSpace(techno)
				}
			}
		}

		if projectPathErr.IsError() {
			println(projectPathQuestion)
			line, rlErr := projectPathReader.Readline()

			if rlErr != nil {
				if err != "" {
					err += CharBack() + Tab()
				} else {
					err += "ERROR : "
				}
				err += projectPathErr.Error()
			} else {
				if strings.TrimSpace(line) == "" {
					if err != "" {
						err += CharBack() + Tab()
					} else {
						err += "ERROR : "
					}
					err += projectPathErr.Error()
				} else {
					var pwd, _ = os.Getwd()
					projectPath = line
					for _, el := range []string{"$(pwd)", "${pwd}"} {
						projectPath = strings.Replace(projectPath, el, pwd, -1)
					}
					projectPath = strings.Replace(projectPath, " "+Slash()+" ", Slash(), -1)
					projectPath = strings.TrimSpace(projectPath)
				}
			}
		}
	}

	if strings.TrimSpace(err) == "" {
		cb(techno, strings.ReplaceAll(projectPath, Slash()+Slash(), Slash()))
	} else {
		color.Red.Println(err)
	}
}

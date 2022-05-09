package main

import (
	"github.com/gookit/color"
	"go.pkg/nchoquet/config_files"
	"go.pkg/nchoquet/files"
	. "go.pkg/nchoquet/helpers"
	"go.pkg/nchoquet/run"
)

func main() {
	if IsBuild() {
		run.OnRun(func(techno string, projectPath string) {
			var err error
			for path, content := range config_files.ConfigFiles[techno] {
				err = files.Create(projectPath+Slash()+path, content)
				if err != nil {
					color.Red.Println(err.Error())
					break
				}
			}

			if err == nil {
				color.Green.Println("le projet " + techno + " à bien été généré dans le répertoire " + projectPath + " !")
			}
		})
	}
}

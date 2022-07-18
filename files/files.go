package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"test_go_webserver/configFiles"
	. "test_go_webserver/helpers"
	"test_go_webserver/pages/helpers"
	"test_go_webserver/technos"
)

func Exists(path string) bool {
	path = strings.ReplaceAll(path, Slash()+Slash(), Slash())

	if strings.Contains(path, ".") {
		_, err := ioutil.ReadFile(path)
		if err != nil {
			return false
		}
		return true
	} else {
		_, err := ioutil.ReadDir(path)
		if err != nil {
			return false
		}
		return true
	}
}

func Create(path string, content string) error {
	path = strings.ReplaceAll(path, Slash()+Slash(), Slash())

	var splitPath = strings.Split(path, Slash())

	_, _ = ArrayPop(&splitPath)

	var _path = strings.Join(splitPath, Slash())

	var dErr = os.MkdirAll(_path, 0777)
	if dErr != nil {
		return fmt.Errorf("%s directory create error", dErr.Error())
	}

	if Exists(path) {
		return fmt.Errorf("%s file already exists", path)
	}

	var f, cErr = os.Create(path)
	if cErr != nil {
		return fmt.Errorf("%s file open error", cErr.Error())
	}

	var _, wErr = f.Write([]byte(content))
	if wErr != nil {
		return fmt.Errorf("%s file write error", wErr.Error())
	}

	var err = f.Close()
	if err != nil {
		return fmt.Errorf("%s file close error", err.Error())
	}

	return nil
}

func ProjectGeneration(projectPath string, techno technos.Techno, projectName *string) (alert Alert) {
	var err error
	technoName := techno.Name
	technoValue := techno.Value

	for path, content := range configFiles.ConfigFiles[technoName] {
		path, err = helpers.ParseString("path", path, map[string]interface{}{
			"ProjectName": &projectName,
		})
		if err != nil {
			alert.Message = fmt.Sprintf("erreur template path: %s", err.Error())
			alert.Type = ERROR
			return alert
		}

		content, err = helpers.ParseString("content", content, map[string]interface{}{
			"ProjectName": &projectName,
		})
		if err != nil {
			alert.Message = fmt.Sprintf("erreur template content: %s", err.Error())
			alert.Type = ERROR
			return alert
		}

		err = Create(projectPath+Slash()+path, content)
		if err != nil {
			alert.Message = fmt.Sprintf("erreur: %s", err.Error())
			alert.Type = ERROR
			println(err.Error())
			return alert
		}
	}

	if err == nil && Exists(projectPath) {
		alert.Message = fmt.Sprintf("Le projet %s à bien été généré dans le répertoire %s !", technoName, projectPath)
		alert.Type = SUCCESS
		println("Le projet " + technoName + " à bien été généré dans le répertoire " + projectPath + " !")
	} else {
		alert.Message = fmt.Sprintf("Une erreur est survenue lors de la génération du projet %s dans le répertoire %s !", technoValue, projectPath)
		alert.Type = ERROR
		println("Une erreur est survenue lors de la génération du projet " + technoValue + " dans le répertoire " + projectPath + " !")
	}

	return alert
}

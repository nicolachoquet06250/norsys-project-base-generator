package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"test_go_webserver/configFiles"
	. "test_go_webserver/helpers"
	"test_go_webserver/technos"
)

func arrayPop[T any](s *[]T) (v T, err error) {
	if len(*s) == 0 {
		var s T
		return s, fmt.Errorf("can't remove last element of array")
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e, nil
}

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

	_, _ = arrayPop(&splitPath)

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

func ProjectGeneration(projectPath string, techno technos.Techno) {
	var err error
	technoName := techno.Name
	technoValue := techno.Value

	for path, content := range configFiles.ConfigFiles[technoName] {
		err = Create(projectPath+Slash()+path, content)

		if err != nil {
			println(err.Error())
			break
		}
	}

	if err == nil && Exists(projectPath) {
		println("Le projet " + technoName + " à bien été généré dans le répertoire " + projectPath + " !")
	} else {
		println("Une erreur est survenue lors de la génération du projet " + technoValue + " dans le répertoire " + projectPath + " !")
	}
}

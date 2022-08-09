//go:build windows && !linux && !freebsd && !netbsd && !openbsd && !darwin && !js
// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package notifications

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"io/ioutil"
	"npbg/notify"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func generateTempScript(content string) (file string, err error) {
	id, _ := uuid.NewV4()
	file = filepath.Join(os.TempDir(), id.String()+".ps1")
	println("script file : ", file)
	bomUtf8 := []byte{0xEF, 0xBB, 0xBF}
	out := append(bomUtf8, []byte(content)...)
	err = ioutil.WriteFile(file, out, 0600)
	if err != nil {
		file = ""
		return
	}
	return
}

func removeTempScript(file string) error {
	if file != "" {
		return os.Remove(file)
	}
	return nil
}

func compatibleWithXmlToast() bool {
	var err error
	var file string

	file, err = generateTempScript("(Get-PSSessionConfiguration -Name Test).LanguageMode")
	defer removeTempScript(file)
	if err != nil {
		if err != nil {
			fmt.Println(fmt.Errorf("%s", err))
			return false
		}
	}

	cmd := exec.Command("PowerShell", "-ExecutionPolicy", "Bypass", "-File", file)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	var out []byte
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Errorf("%s", err))
		return false
	}

	if string(out) != "FullLanguage" {
		return false
	}

	return true
}

func Notify(title, message, appIcon string, appId *string) error {
	if compatibleWithXmlToast() {
		return beeep.Notify(title, message, appIcon, appId)
	}
	return fmt.Errorf("your computer is not compatible with notifications")
}

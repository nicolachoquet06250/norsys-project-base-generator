//go:build linux || freebsd || netbsd || openbsd
// +build linux freebsd netbsd openbsd

package notifications

func Notify(title, message, appIcon string, appId *string) error {
	return beeep.Notify(title, message, appIcon, appId)
}

//go:build linux || freebsd || netbsd || openbsd
// +build linux freebsd netbsd openbsd

package notifications

import (
	"npbg/notify"
)

func Notify(title, message, appIcon string, appId *string) error {
	return beeep.Notify(title, message, appIcon, appId)
}

//go:build (linux && nodbus) || (freebsd && nodbus) || (netbsd && nodbus) || (openbsd && nodbus)
// +build linux,nodbus freebsd,nodbus netbsd,nodbus openbsd,nodbus

package notifications

import (
	"npbg/notify"
)

func Notify(title, message, appIcon string, appId *string) error {
	return beeep.Notify(title, message, appIcon, appId)
}

//go:build darwin && !linux && !freebsd && !netbsd && !openbsd && !windows && !js
// +build darwin,!linux,!freebsd,!netbsd,!openbsd,!windows,!js

package notifications

import (
	"npbg/notify"
)

func Notify(title, message, appIcon string, appId *string) error {
	return beeep.Notify(title, message, appIcon, appId)
}

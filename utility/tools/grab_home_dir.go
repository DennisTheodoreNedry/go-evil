package tools

import (
	"os/user"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Grabs the current users home directory
func Grab_home_dir() string {
	path, err := user.Current()

	if err != nil {
		notify.Error(err.Error(), "tools.Grab_home_dir()")
	}

	return path.HomeDir
}

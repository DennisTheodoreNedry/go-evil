package tools

import (
	"os"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Grabs the current working path
func Grab_CWD() string {
	path, err := os.Getwd()
	if err != nil {
		notify.Error(err.Error(), "tools.Grab_CWD()")
	}
	return path
}

package tools

import (
	"os/user"
	"strings"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Returns the username of the current user
func Grab_username() string {
	user, err := user.Current()

	if err != nil {
		notify.Error(err.Error(), "tools.Grab_username()")
	}

	to_return := user.Username

	if strings.Contains(to_return, "\\") { // Only occurs on windows so far
		split := strings.Split(to_return, "\\")
		to_return = split[1]
	}

	return to_return
}

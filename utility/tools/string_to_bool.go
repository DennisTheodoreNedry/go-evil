package tools

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Converts the provided string to into a boolean
func String_to_boolean(value string) bool {
	to_return := false

	if value != "true" && value != "false" {
		notify.Error(fmt.Sprintf("Needed true/false, recieved %s", value), "tools.String_to_boolean()")
	}

	if value == "true" {
		to_return = true
	}

	return to_return
}

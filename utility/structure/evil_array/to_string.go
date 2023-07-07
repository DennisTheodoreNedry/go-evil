package evilarray

import (
	"fmt"
	"strings"

	notify "github.com/s9rA16Bf4/notify_handler"
)

// Returns the contents of the array as a string
// Possible formats are:
//   - evil - Format ${...}$
//   - array - Format []string{...}
func (object *Evil_array_t) To_string(format string) string {
	header := ""
	footer := ""

	format = strings.ToLower(format) // Makes it lowercase

	switch format {
	case "evil":
		header = "${"
		footer = "}$"
	case "array":
		header = "[]string{"
		footer = "}"
	default:
		notify.Error(fmt.Sprintf("Unknown format %s", format), "evil_array.To_string()", 1)
	}

	toReturn := header

	for _, cont := range object.gut {
		toReturn += fmt.Sprintf("\"%s\",", cont)
	}
	toReturn += footer

	return toReturn
}

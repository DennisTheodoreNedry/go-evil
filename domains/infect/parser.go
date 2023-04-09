package infect

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/infect/path"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// The main parser for the Infect domain
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "path":
		call, s_json = path.Path(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "infect.Parser()")

	}

	return call, s_json
}

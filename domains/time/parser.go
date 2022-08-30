package time

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// The main parser for the time domain
//
//
func Parser(function string, value string, s_json string) (string, string) {
	call := ""

	switch function {
	case "until":
		call, s_json = Until(s_json, value)

	case "sleep":
		call, s_json = Sleep(s_json, value)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call, s_json
}

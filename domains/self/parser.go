package self

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// The main parser for the self domain
//
//
func Parser(function string, value string, s_json string) (string, string) {
	call := ""

	switch function {
	case "call":
		call, s_json = Call_function(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call, s_json
}

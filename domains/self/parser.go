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
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "call":
		call, s_json = Call_function(value, s_json)

	case "include":
		s_json = Include(value, s_json)

	case "set_run":
		call, s_json = Set(false, value, s_json)

	case "set_comp":
		call, s_json = Set(true, value, s_json)

	case "add_random_var":
		s_json = Add_random_variable(value, s_json)

	case "add_random_func":
		call, s_json = Add_random_function(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call, s_json
}

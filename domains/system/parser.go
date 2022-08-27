package system

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// The main parser for the system domain
//
//
func Parser(function string, value string, s_json string) (string, string) {
	call := ""

	switch function {
	case "out":
		call, s_json = Out(s_json, value)

	case "exit":
		call, s_json = Exit(s_json, value)

	// case "exec":
	// 	to_return = append(to_return, Exec(value)...)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call, s_json
}

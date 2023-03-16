package network

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
	case "ping":
		call, s_json = ping(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "network.Parser()")

	}

	return call, s_json
}

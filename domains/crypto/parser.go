package crypto

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// The main parser for the random domain
//
//
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "encrypt":
		s_json = encrypt(value, s_json)

	case "set_method":
		s_json = set_method(value, s_json)

	case "set_key":
		s_json = set_key(value, s_json)

	case "set_target":
		s_json = set_target(value, s_json)

	case "decrypt":
		s_json = decrypt(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call, s_json
}

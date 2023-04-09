package random

import (
	"fmt"

	evil_int "github.com/TeamPhoneix/go-evil/domains/random/int"
	evil_string "github.com/TeamPhoneix/go-evil/domains/random/string"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// The main parser for the random domain
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "int":
		s_json = evil_int.Generate(value, s_json)

	case "string":
		s_json = evil_string.Generate(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call, s_json
}

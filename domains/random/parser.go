package random

import (
	"fmt"

	evil_int "github.com/s9rA16Bf4/go-evil/domains/random/int"
	evil_string "github.com/s9rA16Bf4/go-evil/domains/random/string"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "int":
		evil_int.Generate(value, data_object)

	case "string":
		evil_string.Generate(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call
}

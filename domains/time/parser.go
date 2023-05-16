package time

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/time/sleep"
	"github.com/s9rA16Bf4/go-evil/domains/time/until"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "until":
		call = until.Until(value, data_object)

	case "sleep":
		call = sleep.Sleep(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call
}

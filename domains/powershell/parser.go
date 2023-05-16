package powershell

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/powershell/policy"

	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	if data_object.Target_os != "windows" {
		notify.Error("The target OS must be 'windows' to be able to use the 'powershell' domain!", "powershell.Parser()")
	}

	switch function {
	case "set_execution_policy":
		call = policy.Set_execution(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "powershell.Parser()")

	}

	return call
}

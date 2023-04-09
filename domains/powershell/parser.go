package powershell

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/powershell/policy"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// The main parser for the Powershell domain
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	data_object := structure.Receive(s_json)

	if data_object.Target_os != "windows" {
		notify.Error("The target OS must be 'windows' to be able to use the 'powershell' domain!", "powershell.Parser()")
	}

	switch function {
	case "set_execution_policy":
		call, s_json = policy.Set_execution(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "powershell.Parser()")

	}

	return call, s_json
}

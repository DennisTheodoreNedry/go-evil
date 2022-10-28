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
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "out":
		call, s_json = Out(s_json, value)

	case "outln":
		call, s_json = Outln(s_json, value)

	case "exit":
		call, s_json = Exit(s_json, value)

	case "exec":
		call, s_json = Exec(s_json, value)

	case "abort":
		call, s_json = Abort(s_json, value)

	case "reboot":
		call, s_json = Reboot(s_json)

	case "shutdown":
		call, s_json = Shutdown(s_json)

	case "add_to_startup":
		call, s_json = Add_to_startup(s_json)

	case "write":
		call, s_json = write(s_json, value)

	case "read":
		call, s_json = read(s_json, value)

	case "list_dir":
		call, s_json = list_dir(s_json, value)

	case "in":
		call, s_json = input(s_json)

	case "remove":
		call, s_json = remove(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call, s_json
}

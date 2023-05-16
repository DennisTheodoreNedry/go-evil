package interfaces

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Grabs the current wireless interface
// Input None
// The result will be the interface name and mac adress which are placed into seperate runtime variables
func Get_interface(value string, data_object *json.Json_t) []string {
	function_call := "get_interface"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{""},
		Gut: []string{
			"i_name, i_mac := coldfire.Iface()",
			"spine.variable.set(i_name)",
			"spine.variable.set(i_mac)",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}
}

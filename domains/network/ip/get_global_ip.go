package ip

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Get the global ip address
// The result is placed in a runtime variable
func Get_global(value string, data_object *json.Json_t) []string {
	function_call := "get_global_ip"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{""},
		Gut: []string{
			"spine.variable.set(coldfire.GetGlobalIp())",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}
}

package interfaces

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Grabs the current wireless interface
// Input None
// The result will be the interface name and mac adress which are placed into seperate runtime variables
func Get_interface(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_interface"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{""},
		Gut: []string{
			"i_name, i_mac := coldfire.Iface()",
			"spine.variable.set(i_name)",
			"spine.variable.set(i_mac)",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

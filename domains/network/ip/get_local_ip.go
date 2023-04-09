package ip

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Get the local ip address
// The result is placed in a runtime variable
func Get_local(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_local_ip"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"spine.variable.set(coldfire.GetLocalIp())",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

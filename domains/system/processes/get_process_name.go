package processes

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Tries to grab all process names
// Input None
// The return is an evil array containing all found users which is placed in a runtime variable
func Get_names(value string, data_object *json.Json_t) []string {
	function_call := "get_processes_names"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{""},
		Gut: []string{
			"processes, err := coldfire.Processes()",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
			"arr := structure.Create_evil_object(\"\")",
			"for _, value := range processes{",
			"arr.Append(value)",
			"}",
			"spine.variable.set(arr.To_string(\"evil\"))",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/s9rA16Bf4/go-evil/utility/structure")

	return []string{fmt.Sprintf("%s()", function_call)}
}

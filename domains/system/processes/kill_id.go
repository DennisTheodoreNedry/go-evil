package processes

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Tries to terminate a process based on it's pid
// Input must therefore be the pid to utilize
func Kill_id(value string, data_object *json.Json_t) []string {
	function_call := "kill_process_id"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"value1 := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_1)))",
			"err := coldfire.PkillPid(value1)",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
		}})

	data_object.Add_go_import("github.com/s9rA16Bf4/go-evil/utility/tools")
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}
}

package stop

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Exits the malware
func Exit(return_code string, data_object *json.Json_t) []string {

	function_call := "Exit"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"lvl := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"spine.return_code = tools.String_to_int(lvl)",
			"spine.terminate = true",
		}})

	data_object.Add_go_import("github.com/s9rA16Bf4/go-evil/utility/tools")

	parameter_1 := data_object.Generate_int_array_parameter(return_code)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}
}

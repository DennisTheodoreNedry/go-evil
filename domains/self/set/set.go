package set

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
)

// Sets a compiletime/runtime variable with a value
func Set(compile_time bool, value string, data_object *json.Json_t) []string {
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	if compile_time {
		data_object.Set_variable_value(value)
		return []string{}

	} else {
		function_call := "set_runtime"

		data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
			Parameters: []string{"repr_1 []int"},

			Gut: []string{
				"value := spine.variable.get(spine.alpha.construct_string(repr_1))",
				"spine.variable.set(value)",
			}})

		parameter_1 := data_object.Generate_int_array_parameter(value)

		return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}
	}
}

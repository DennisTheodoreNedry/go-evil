package set

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Sets a compiletime/runtime variable with a value
func Set(compile_time bool, value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	if compile_time {
		data_object.Set_variable_value(value)
		return []string{}, structure.Send(data_object)

	} else {
		function_call := "set_runtime"

		data_object.Add_go_function([]string{
			fmt.Sprintf("func %s(repr_1 []int){", function_call),
			"value := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"spine.variable.set(value)",
			"}",
		})

		parameter_1 := data_object.Generate_int_array_parameter(value)

		return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
	}
}

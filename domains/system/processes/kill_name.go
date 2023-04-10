package processes

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Tries to terminate a process based on it's name
// Input must therefore be the name to utilize
func Kill_name(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "kill_process_name"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"value1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"err := coldfire.PkillName(value1)",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

package io

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Removes the target file and folder if they are empty
func Remove(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "remove"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1  []int"},
		Gut: []string{
			"target := spine.alpha.construct_string(repr)",
			"target = spine.variable.get(target)",
			"err := os.Remove(target)",
			"if err != nil{",
			"spine.log(err.Error())",
			"return",
			"}",
		}})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)

}

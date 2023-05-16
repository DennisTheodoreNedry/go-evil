package io

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Prints a message to the screen
func Out(msg string, data_object *json.Json_t) []string {
	function_call := "Out"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1  []int"},
		Gut: []string{
			"s_msg := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"fmt.Print(s_msg)",
		}})

	data_object.Add_go_import("fmt")

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(msg)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}
}

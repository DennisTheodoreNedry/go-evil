package base64

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

//
//
// Encodes the provided string and places the result in a runtime variable
//
//
func encode(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "base64_encode"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"value1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"spine.variable.set(coldfire.B64E(value1))",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

//
//
// Decodes the provided string and places the result in a runtime variable
//
//
func decode(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "base64_decode"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"value1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"spine.variable.set(coldfire.B64D(value1))",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

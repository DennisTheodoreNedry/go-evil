package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Prints a message to the screen
func Out(s_json string, msg string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Out"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(msg []int){", function_call),
		"s_msg := spine.variable.get(spine.alpha.construct_string(msg))",
		"fmt.Print(s_msg)",
		"}"})

	data_object.Add_go_import("fmt")

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(msg)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)
}

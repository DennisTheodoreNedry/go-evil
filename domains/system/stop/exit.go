package stop

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Exits the malware
func Exit(s_json string, return_code string) ([]string, string) {
	data_object := structure.Receive(s_json)

	function_call := "Exit"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"lvl := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"value := tools.String_to_int(lvl)",
		"os.Exit(value)",

		"}"})

	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("os")

	parameter_1 := data_object.Generate_int_array_parameter(return_code)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

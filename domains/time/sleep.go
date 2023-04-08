package time

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Makes the malware sleep for an n amount of seconds
func Sleep(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Sleep"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"i_value := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_1)))",
		"time.Sleep(time.Duration(i_value) * time.Second)",
		"}",
	})

	data_object.Add_go_import("time")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

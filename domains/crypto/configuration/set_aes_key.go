package configuration

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Sets the aes key used for encrypting
func Set_aes_key(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	function_call := "set_aes_key"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"key := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"spine.crypt.set_aes_key(key)",
		}})

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)
}

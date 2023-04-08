package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Reads the contents of a file and places the result into a runtime variable
func read(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "read"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr []int){", function_call),
		"path := spine.alpha.construct_string(repr)",
		"path = spine.variable.get(path)",
		"gut, err := ioutil.ReadFile(path)",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"spine.variable.set(string(gut))",
		"}"})

	data_object.Add_go_import("io/ioutil")

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)

}

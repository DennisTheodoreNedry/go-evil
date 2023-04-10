package call

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Adds a function function_call to the src code
// Calls function of the type 'c'
func Function(func_name string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := fmt.Sprintf("function_call_%s", tools.Generate_random_n_string(16))

	func_name = tools.Erase_delimiter(func_name, []string{"\""}, -1) // Removes all " from the string

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut:        []string{fmt.Sprintf("%s()", func_name)}})

	return []string{function_call}, structure.Send(data_object)
}

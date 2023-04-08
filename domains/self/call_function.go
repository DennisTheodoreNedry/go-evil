package self

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Adds a function function_call to the src code
// Calls function of the type 'c'
func Call_function(func_name string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := []string{"function_call()"}

	func_name = tools.Erase_delimiter(func_name, []string{"\""}, -1) // Removes all " from the string

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call[0]),
		fmt.Sprintf("%s()", func_name),
		"}",
	})

	return function_call, structure.Send(data_object)
}

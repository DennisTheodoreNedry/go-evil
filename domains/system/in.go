package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Takes a user input and saves the result in a runtime variable
func input(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "input"
	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"var input string",
		"fmt.Scanln(&input)",
		"spine.variable.set(input)",
		"}"})

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)

}

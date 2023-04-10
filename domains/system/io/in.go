package io

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Takes a user input and saves the result in a runtime variable
func Input(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "input"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut: []string{
			"var input string",
			"fmt.Scanln(&input)",
			"spine.variable.set(input)",
		}})

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)

}

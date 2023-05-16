package io

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Takes a user input and saves the result in a runtime variable
func Input(data_object *json.Json_t) []string {
	function_call := "input"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut: []string{
			"var input string",
			"fmt.Scanln(&input)",
			"spine.variable.set(input)",
		}})

	return []string{fmt.Sprintf("%s()", function_call)}

}

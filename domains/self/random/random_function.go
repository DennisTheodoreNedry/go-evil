package random

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/self/random/generate"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Adds a random function to the source code
func Add_function(amount string, data_object *json.Json_t) []string {
	amount = tools.Erase_delimiter(amount, []string{"\""}, -1)
	calls := []string{}

	i_value := tools.String_to_int(amount)
	if i_value == -1 {
		notify.Error(fmt.Sprintf("Unknown amount '%d'", i_value), "self.Add_random_variable()")
	}

	for i := 0; i < i_value; i++ {

		// Generate the function name
		function_name := generate.Generate_function_name()

		// Generate the return type
		return_type := generate.Generate_return_type()

		// Generate a random amount of parameters
		parameter_array, sending_values := generate.Generate_parameters()

		calls = append(calls, fmt.Sprintf("%s(%s)", function_name, sending_values))

		// Construct the function body
		body := generate.Generate_function_body(return_type)

		// Add it
		data_object.Add_go_function(functions.Go_func_t{Name: function_name, Func_type: "", Part_of_struct: "", Return_type: return_type,
			Parameters: parameter_array, Gut: body})
	}

	data_object.Add_go_import("fmt")

	return calls
}

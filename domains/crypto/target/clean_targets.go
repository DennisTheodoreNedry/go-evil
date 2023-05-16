package target

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Removes every previously added target
func Clean(value string, data_object *json.Json_t) []string {
	function_call := "clean_targets"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut: []string{
			"spine.crypt.target = []string{}",
		}})

	return []string{fmt.Sprintf("%s()", function_call)}
}

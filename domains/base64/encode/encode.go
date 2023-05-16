package encode

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Encodes the provided string and places the result in a runtime variable
func Encode(value string, data_object *json.Json_t) []string {
	function_call := "base64_encode"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"value1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"spine.variable.set(coldfire.B64E(value1))",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}
}

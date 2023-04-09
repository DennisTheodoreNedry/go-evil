package dns

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Peforms a dns lookup
func Lookup(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "dns_lookup"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"param_1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"result, err := coldfire.DnsLookup(param_1)",
		"if err != nil {",
		"spine.log(err.Error())",
		"}",
		"spine.variable.set(strings.Join(result,\" \"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("strings")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

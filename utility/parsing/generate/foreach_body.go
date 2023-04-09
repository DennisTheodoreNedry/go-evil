package generate

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

var foreach_call = 0

// Construcs the code needed for a "foreach" loop
func Construct_foreach_loop(condition string, body []string, s_json string) ([]string, string) {
	function_call := []string{fmt.Sprintf("foreach_%d", foreach_call)}
	foreach_call++

	body_calls, s_json := Generate_golang_code(body, s_json) // Converts the code for the foreach body
	data_object := structure.Receive(s_json)
	arr := structure.Create_evil_object(condition)

	final_body := []string{fmt.Sprintf(
		"func %s(values []string){", function_call[0]),

		"for i, value := range values{",
		"value = spine.variable.get(value)",
		"arr := structure.Create_evil_object(value)",
		"if arr.Length() > 0 {",
		"result := arr.Dump()",
		"values[i] = result[0]",
		"values = append(values, result[1:]...)",
		"}}",
	}

	final_body = append(final_body, "for _, value := range values{")
	final_body = append(final_body, "spine.variable.foreach = value")

	final_body = append(final_body, body_calls...)
	final_body = append(final_body, "}}")

	data_object.Add_go_function(final_body)

	if arr.Length() != 0 {
		function_call[0] = fmt.Sprintf("%s(%s)", function_call[0], arr.To_string("array"))
	} else {
		function_call[0] = fmt.Sprintf("%s([]string{%s})", function_call[0], condition)
	}

	data_object.Add_go_import("fmt")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return function_call, structure.Send(data_object)
}

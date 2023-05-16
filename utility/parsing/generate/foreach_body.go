package generate

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

var foreach_call = 0

// Construcs the code needed for a "foreach" loop
func Construct_foreach_loop(condition string, body []string, data_object *json.Json_t) []string {
	function_call := []string{fmt.Sprintf("foreach_%d", foreach_call)}
	foreach_call++

	body_calls := Generate_golang_code(body, data_object) // Converts the code for the foreach body
	arr := structure.Create_evil_object(condition)

	final_body := []string{
		"for i, value := range values{",
		"value = spine.variable.get(value)",
		"arr := structure.Create_evil_object(value)",
		"if arr.Length() > 0 {",
		"result := arr.Dump()",
		"values[i] = result[0]",
		"values = append(values, result[1:]...)",
		"}}",
		"for _, value := range values{",
		"spine.variable.foreach = value",
	}

	final_body = append(final_body, body_calls...)
	final_body = append(final_body, "}")

	data_object.Add_go_function(functions.Go_func_t{Name: function_call[0], Func_type: "", Part_of_struct: "",
		Return_type: "", Parameters: []string{"values []string"}, Gut: final_body})

	if arr.Length() != 0 {
		function_call[0] = fmt.Sprintf("%s(%s)", function_call[0], arr.To_string("array"))
	} else {
		function_call[0] = fmt.Sprintf("%s([]string{%s})", function_call[0], condition)
	}

	data_object.Add_go_import("fmt")
	data_object.Add_go_import("github.com/s9rA16Bf4/go-evil/utility/structure")

	return function_call
}

package parsing

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

//
//
// Construcs the code needed for a "foreach" loop
//
//
func construct_foreach_loop(condition string, body []string, s_json string) ([]string, string) {
	call := []string{"foreach"}

	body_calls, s_json := convert_code(body, s_json) // Converts the code for the foreach body
	data_object := structure.Receive(s_json)
	arr := structure.Create_evil_object(condition)

	final_body := []string{fmt.Sprintf(
		"func %s(values []string){", call[0]),

		"for _, value := range values{",
		"value = spine.variable.get(value)",
		"arr := structure.Create_evil_object(value)",
		"if arr.Length() > 0 {",
		"for _, sub_value := range arr.Dump(){",
		"spine.variable.foreach = sub_value",
	}

	final_body = append(final_body, body_calls...)
	final_body = append(final_body, "}")

	final_body = append(final_body, "}else{", "spine.variable.foreach = value")

	final_body = append(final_body, body_calls...)
	final_body = append(final_body, "}}}")

	data_object.Add_go_function(final_body)

	if arr.Length() != 0 {
		call[0] = fmt.Sprintf("%s(%s)", call[0], arr.To_string("array"))
	} else {
		call[0] = fmt.Sprintf("%s([]string{%s})", call[0], condition)
	}

	data_object.Add_go_import("fmt")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return call, structure.Send(data_object)
}

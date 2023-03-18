package parsing

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

var foreach_call = 0

//
//
// Construcs the code needed for a "foreach" loop
//
//
func construct_foreach_loop(condition string, body []string, s_json string) ([]string, string) {
	call := []string{fmt.Sprintf("foreach_%d", foreach_call)}
	foreach_call++

	body_calls, s_json := generate_body_code(body, s_json) // Converts the code for the foreach body
	data_object := structure.Receive(s_json)
	arr := structure.Create_evil_object(condition)

	final_body := []string{fmt.Sprintf(
		"func %s(values []string){", call[0]),

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
		call[0] = fmt.Sprintf("%s(%s)", call[0], arr.To_string("array"))
	} else {
		call[0] = fmt.Sprintf("%s([]string{%s})", call[0], condition)
	}

	data_object.Add_go_import("fmt")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return call, structure.Send(data_object)
}

//
//
// Gathers all data needed for an foreach statement
//
//
func get_foreach_body(index *int, gut []string) []string {
	body := []string{}
	*index++ // Skips the header which is important as we otherwise get stuck in an endless loop

	for ; *index < len(gut); *index++ { // Grabs all data between the header and footer, but also fast forwards the index
		footer := tools.Contains(gut[*index], []string{GET_FOREACH_FOOTER})
		footer_reached := footer[GET_FOREACH_FOOTER]

		if footer_reached {
			break
		} else {
			body = append(body, gut[*index])
		}
	}

	return body
}

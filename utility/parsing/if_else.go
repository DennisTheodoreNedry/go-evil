package parsing

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

var if_else_call = 0

//
//
// Construcs the code needed for a "if/else" statement
//
//
func construct_if_else(condition string, if_true_body []string, if_false_body []string, s_json string) ([]string, string) {
	function_call := fmt.Sprintf("if_%d", if_else_call)
	if_else_call++

	if_true_body_calls, s_json := generate_body_code(if_true_body, s_json)   // Converts the if true code
	if_false_body_calls, s_json := generate_body_code(if_false_body, s_json) // Converts the else code

	data_object := structure.Receive(s_json)
	arr := structure.Create_evil_object(condition)

	if arr.Length() != 3 {
		notify.Error(fmt.Sprintf("Expected three values, but recieved %d", arr.Length()), "if_else.construct_if_else()")
	}

	compare_operator := arr.Get(1)
	switch compare_operator {
	case ">", "<", "==", "!=", "<=", ">=":
	default:
		notify.Error(fmt.Sprintf("Unknown and/or illegal operator %s", compare_operator), "if_else.construct_if_else()")
	}

	final_body := []string{fmt.Sprintf(
		"func %s(repr_value1 []int, repr_value2 []int){", function_call),
		"value1 := spine.variable.get(spine.alpha.construct_string(repr_value1))",
		"value2 := spine.variable.get(spine.alpha.construct_string(repr_value2))",
		fmt.Sprintf("if (value1 %s value2){", compare_operator),
	}

	final_body = append(final_body, if_true_body_calls...)
	final_body = append(final_body, "}else{")
	final_body = append(final_body, if_false_body_calls...)
	final_body = append(final_body, "}}")

	data_object.Add_go_function(final_body)

	data_object.Add_go_import("fmt")

	parameter_1 := data_object.Generate_int_array_parameter(arr.Get(0))
	parameter_2 := data_object.Generate_int_array_parameter(arr.Get(2))

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, parameter_1, parameter_2)}, structure.Send(data_object)
}

//
//
// Gathers all data needed for an if/else statement
//
//
func get_if_else_body(index *int, gut []string) ([]string, []string) {
	if_true_body := []string{}
	if_false_body := []string{}

	*index++ // Skips the header which is important as we otherwise get stuck in an endless loop

	else_statement := map[string]bool{}
	reached_else := false

	// Grab the if true and if false body
	for ; *index < len(gut); *index++ { // Grabs all data between the header and footer, but also fast forwards the index
		footer := tools.Contains(gut[*index], []string{GET_IF_ELSE_FOOTER})
		footer_reached := footer[GET_IF_ELSE_FOOTER]

		if !reached_else { // Only do this once
			else_statement = tools.Contains(gut[*index], []string{GET_ELSE_HEADER})
			reached_else = else_statement[GET_ELSE_HEADER]
		}

		if footer_reached {
			break
		} else if !reached_else {
			if_true_body = append(if_true_body, gut[*index])
		} else {
			if_false_body = append(if_false_body, gut[*index])
		}

	}

	return if_true_body, if_false_body
}

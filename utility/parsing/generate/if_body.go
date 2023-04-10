package generate

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

var if_else_call = 0

// Construcs the code needed for a "if/else" statement
func Construct_if_else(condition string, if_true_body []string, if_false_body []string, s_json string) ([]string, string) {
	function_call := fmt.Sprintf("if_%d", if_else_call)
	if_else_call++

	if_true_body_calls, s_json := Generate_golang_code(if_true_body, s_json)   // Converts the if true code
	if_false_body_calls, s_json := Generate_golang_code(if_false_body, s_json) // Converts the else code

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

	body := []string{
		"value1 := spine.variable.get(spine.alpha.construct_string(repr_value1))",
		"value2 := spine.variable.get(spine.alpha.construct_string(repr_value2))",
		fmt.Sprintf("if (value1 %s value2){", compare_operator),
	}

	body = append(body, if_true_body_calls...)
	body = append(body, "}else{")
	body = append(body, if_false_body_calls...)
	body = append(body, "}")

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "",
		Return_type: "", Parameters: []string{"repr_value1 []int", "repr_value2 []int"}, Gut: body})

	data_object.Add_go_import("fmt")

	parameter_1 := data_object.Generate_int_array_parameter(arr.Get(0))
	parameter_2 := data_object.Generate_int_array_parameter(arr.Get(2))

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, parameter_1, parameter_2)}, structure.Send(data_object)
}

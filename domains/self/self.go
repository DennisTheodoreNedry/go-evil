package self

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Adds a function call to the src code
//
//
func Call_function(func_name string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	call := []string{"call()"}

	func_name = tools.Erase_delimiter(func_name, []string{"\""}, -1) // Removes all " from the string

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", call[0]),
		fmt.Sprintf("%s()", func_name),
		"}",
	})

	return call, structure.Send(data_object)
}

//
//
// Includes the provided file in the malware, the result can be found in one of the compiler time variables
//
//
func Include(file_path string, s_json string) string {
	data_object := structure.Receive(s_json)
	file_path = tools.Erase_delimiter(file_path, []string{"\""}, -1)

	file_gut, err := ioutil.ReadFile(file_path)
	if err != nil {
		notify.Error(err.Error(), "self.Include()")
	}
	new_const := tools.Generate_random_string()
	final_line := fmt.Sprintf("var %s = \"[HEX];,", new_const)

	for _, line := range file_gut {
		final_line += fmt.Sprintf("%s,", hex.EncodeToString([]byte{line}))
	}

	final_line += "\""

	data_object.Add_go_global(final_line)

	data_object.Set_variable_value(new_const)

	return structure.Send(data_object)
}

//
//
// Sets a compiletime/runtime variable with a value
//
//
func Set(var_type string, value string, s_json string) string {
	data_object := structure.Receive(s_json)
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	switch var_type {
	case "$":
		data_object.Set_variable_value(value)

	case "€":

	default:
		notify.Error(fmt.Sprintf("Unknown variable type '%s'", var_type), "self.Set()")
	}

	return structure.Send(data_object)
}

//
//
// Adds a random variable to the source code
//
//
func Add_random_variable(amount string, s_json string) string {
	data_object := structure.Receive(s_json)

	amount = tools.Erase_delimiter(amount, []string{"\""}, -1)

	i_value := tools.String_to_int(amount)
	if i_value == -1 {
		notify.Error(fmt.Sprintf("Unknown amount '%d'", i_value), "self.Add_random_variable()")
	}

	for i := 0; i < i_value; i++ {
		variable_name := tools.Generate_random_string()
		random_value := tools.Generate_random_string()

		data_object.Add_go_global(fmt.Sprintf("var %s string = \"%s\"", variable_name, random_value))
	}

	return structure.Send(data_object)
}

//
//
// Adds a random function to the source code
//
//
func Add_random_function(amount string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	amount = tools.Erase_delimiter(amount, []string{"\""}, -1)
	calls := []string{}

	i_value := tools.String_to_int(amount)
	if i_value == -1 {
		notify.Error(fmt.Sprintf("Unknown amount '%d'", i_value), "self.Add_random_variable()")
	}

	return_values := []string{"int", "bool", "string", "[]string", "[]int", "[]bool"}
	math_operator := []string{"+", "-", "*", "/", "%"}

	for i := 0; i < i_value; i++ {

		// Generate the function name
		function_name := tools.Generate_random_string()

		// Generate the return type
		return_type := return_values[tools.Generate_random_int_between(0, len(return_values))]

		// Generate a random amount of parameters
		amount_of_parameters := tools.Generate_random_int()
		parameter_line := ""
		sending_values := ""
		for y := 0; y < amount_of_parameters; y++ {
			parameter_line += fmt.Sprintf("param%d string,", y)
			sending_values += fmt.Sprintf("\"%s\",", tools.Generate_random_string())
		}

		calls = append(calls, fmt.Sprintf("%s(%s)", function_name, sending_values))

		// Construct the function body
		body := []string{fmt.Sprintf("func %s(%s) %s{", function_name, parameter_line, return_type)}

		// Generate a certain amount of lines in the function
		body_lines := tools.Generate_random_int_between(1, 32)
		body = append(body, "payload_body := \"A\"")
		body = append(body, "payload_length := 0")

		for line := 0; line < body_lines; line++ {
			line_type := tools.Generate_random_int_between(0, 3)

			switch line_type {
			case 1: // String
				variable := tools.Generate_random_string()
				content := tools.Generate_random_string()

				body = append(body, fmt.Sprintf("%s := \"%s\"", variable, content))
				body = append(body, fmt.Sprintf("payload_body += %s", variable))

			case 2: // Math
				op := math_operator[tools.Generate_random_int_between(0, len(math_operator))]
				a := tools.Generate_random_int()
				b := tools.Generate_random_int()
				c := tools.Generate_random_string()

				body = append(body, fmt.Sprintf("%s := %d %s %d", c, a, op, b))
				body = append(body, fmt.Sprintf("payload_length += %s", c))
			}
		}
		body = append(body, "fmt.Sprintf(\"%s\", payload_body)", "fmt.Sprintf(\"%d\", payload_length)")

		// Generate return value/values
		switch return_type {
		case "int":
			body = append(body, fmt.Sprintf("return %d", tools.Generate_random_int()))

		case "bool":
			body = append(body, fmt.Sprintf("return %t", tools.Generate_random_bool()))

		case "string":
			body = append(body, fmt.Sprintf("return \"%s\"", tools.Generate_random_string()))

		case "[]bool":
			length := tools.Generate_random_int_between(1, 64)
			line := "[]bool{"
			for z := 0; z < length; z++ {
				line += fmt.Sprintf("%t,", tools.Generate_random_bool())
			}
			line += "}"

			body = append(body, fmt.Sprintf("return %s", line))

		case "[]int":
			length := tools.Generate_random_int_between(1, 64)
			line := "[]int{"
			for z := 0; z < length; z++ {
				line += fmt.Sprintf("%d,", tools.Generate_random_int())
			}
			line += "}"

			body = append(body, fmt.Sprintf("return %s", line))

		case "[]string":
			length := tools.Generate_random_int_between(1, 64)
			line := "[]string{"
			for z := 0; z < length; z++ {
				line += fmt.Sprintf("\"%s\",", tools.Generate_random_string())
			}
			line += "}"

			body = append(body, fmt.Sprintf("return %s", line))

		}

		body = append(body, "}")

		// Add it
		data_object.Add_go_function(body)
	}

	data_object.Add_go_import("fmt")

	return calls, structure.Send(data_object)
}

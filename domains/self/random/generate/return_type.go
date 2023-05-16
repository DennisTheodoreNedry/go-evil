package generate

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
)

func Generate_return_type() string {
	return_values := []string{"int", "bool", "string", "[]string", "[]int", "[]bool"}

	return return_values[tools.Generate_random_int_between(0, len(return_values))]
}

func Generate_return_string_line(return_type string) string {
	to_return := ""

	switch return_type {
	case "int":
		to_return = fmt.Sprintf("return %d", tools.Generate_random_int())

	case "bool":
		to_return = fmt.Sprintf("return %t", tools.Generate_random_bool())

	case "string":
		to_return = fmt.Sprintf("return \"%s\"", tools.Generate_random_string())

	case "[]bool":
		length := tools.Generate_random_int_between(1, 64)
		line := "[]bool{"
		for z := 0; z < length; z++ {
			line += fmt.Sprintf("%t,", tools.Generate_random_bool())
		}
		line += "}"

		to_return = fmt.Sprintf("return %s", line)

	case "[]int":
		length := tools.Generate_random_int_between(1, 64)
		line := "[]int{"
		for z := 0; z < length; z++ {
			line += fmt.Sprintf("%d,", tools.Generate_random_int())
		}
		line += "}"

		to_return = fmt.Sprintf("return %s", line)

	case "[]string":
		length := tools.Generate_random_int_between(1, 64)
		line := "[]string{"
		for z := 0; z < length; z++ {
			line += fmt.Sprintf("\"%s\",", tools.Generate_random_string())
		}
		line += "}"

		to_return = fmt.Sprintf("return %s", line)
	}

	return to_return
}

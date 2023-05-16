package generate

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
)

func Generate_math_line() []string {
	body := []string{}
	math_operator := []string{"+", "-", "*", "/", "%"}

	op := math_operator[tools.Generate_random_int_between(0, len(math_operator))]
	a := tools.Generate_random_int()
	b := tools.Generate_random_int()
	c := tools.Generate_random_string()

	body = append(body, fmt.Sprintf("%s := %d %s %d", c, a, op, b))
	body = append(body, fmt.Sprintf("payload_length += %s", c))

	return body
}

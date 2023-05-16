package generate

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
)

func Generate_string_line() []string {
	body := []string{}

	variable := tools.Generate_random_string()
	content := tools.Generate_random_string()

	body = append(body, fmt.Sprintf("%s := \"%s\"", variable, content))
	body = append(body, fmt.Sprintf("payload_body += %s", variable))

	return body
}

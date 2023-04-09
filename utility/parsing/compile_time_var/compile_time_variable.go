package compiletimevar

import (
	"regexp"
	"strings"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Checks the incoming line for compile time variable, and if detected parses it correctly
func Parse_compile_time_vars(value string, s_json string) (string, string) {
	regex := regexp.MustCompile(evil_regex.GET_VAR)
	result := regex.FindAllStringSubmatch(value, -1)

	if len(result) > 0 {
		data_object := structure.Receive(s_json)
		var_call := result[0][1]
		var_id := result[0][2]

		var_value := data_object.Get_var_value(var_id)
		value = strings.ReplaceAll(value, var_call, var_value)

		s_json = structure.Send(data_object)
	}

	return value, s_json
}

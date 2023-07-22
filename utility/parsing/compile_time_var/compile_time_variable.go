package compiletimevar

import (
	"regexp"
	"strings"

	evil_regex "github.com/s9rA16Bf4/go-evil/utility/parsing/regex"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Checks the incoming line for compile time variable, and if detected parses it correctly
func Parse_compile_time_vars(value string, data_object *json.Json_t) string {
	regex := regexp.MustCompile(evil_regex.GET_VAR)
	result := regex.FindAllStringSubmatch(value, -1)

	if len(result) > 0 {
		var_call := result[0][1]
		var_id := result[0][2]

		var_value := data_object.Get_var_value(var_id)
		value = strings.ReplaceAll(value, var_call, var_value)

	}

	return value
}

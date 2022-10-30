package parsing

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TeamPhoneix/go-evil/domains/crypto"
	"github.com/TeamPhoneix/go-evil/domains/random"
	"github.com/TeamPhoneix/go-evil/domains/self"
	"github.com/TeamPhoneix/go-evil/domains/system"
	"github.com/TeamPhoneix/go-evil/domains/time"
	"github.com/TeamPhoneix/go-evil/domains/window"
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Parses the data from the target file and generates function structures from it
//
//
func Build_functions(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(FUNC)
	functions := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(functions) > 0 {
		for _, function := range functions {
			f_type := function[1]
			name := function[2]
			gut := function[3:]

			data_object.Add_function(name, f_type, gut)

		}
	}
	return structure.Send(data_object)
}

//
//
// Converts evil code to golang code and returns it
//
//
func convert_code(gut []string, s_json string) ([]string, string) {
	calls := []string{}

	for i := 0; i < len(gut); i++ {
		line := gut[i]
		call_functions := []string{}

		// Identify which domain to call on
		regex := regexp.MustCompile(DOMAIN_FUNC_VALUE)
		data := regex.FindAllStringSubmatch(line, -1)

		if len(data) > 0 {
			// This makes it easier to figure out what is what
			domain := data[0][1]
			function := data[0][2]
			value := data[0][3]

			call_functions, s_json = grab_code(domain, function, value, s_json)

		} else {
			regex = regexp.MustCompile(GET_FOREACH_HEADER)
			data = regex.FindAllStringSubmatch(line, -1)

			if len(data) > 0 {
				body := []string{}
				i++ // Skips the header

				for ; i < len(gut); i++ { // Grabs all data between the header and footer, but also fast forwards the index
					result := tools.Contains(gut[i], []string{GET_FOREACH_FOOTER})
					status := result[GET_FOREACH_FOOTER]

					if !status {
						body = append(body, gut[i])

					} else { // Footer reached
						break
					}
				}
				call_functions, s_json = construct_foreach_loop(data[0][1], body, s_json)

			}
		}

		if len(call_functions) > 0 { // Don't want any empty lines
			calls = append(calls, call_functions...)
		}
	}

	return calls, s_json
}

//
//
// Grab the contents from each domains parser
//
//
func grab_code(domain string, function string, value string, s_json string) ([]string, string) {
	call_functions := []string{}

	// Translating compile time variables
	regex := regexp.MustCompile(GET_VAR)
	result := regex.FindAllStringSubmatch(value, -1)

	if len(result) > 0 {
		data_object := structure.Receive(s_json)
		var_call := result[0][1]
		Var_type := result[0][2]
		var_id := result[0][3]

		if Var_type == "$" {
			var_value := data_object.Get_var_value(var_id)
			value = strings.ReplaceAll(value, var_call, var_value)
		}

		s_json = structure.Send(data_object)
	}

	// Going through all available domains
	switch domain {
	case "system", "#sys":
		call_functions, s_json = system.Parser(function, value, s_json)

	case "time":
		call_functions, s_json = time.Parser(function, value, s_json)

	case "window":
		call_functions, s_json = window.Parser(function, value, s_json)

	case "self", "#me", "#this":
		call_functions, s_json = self.Parser(function, value, s_json)

	case "random", "#rand":
		call_functions, s_json = random.Parser(function, value, s_json)

	case "crypto":
		call_functions, s_json = crypto.Parser(function, value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown domain '%s'", domain), "functions.grab_code()")
	}

	return call_functions, s_json
}

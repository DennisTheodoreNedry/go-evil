package parsing

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TeamPhoneix/go-evil/domains/self"
	"github.com/TeamPhoneix/go-evil/domains/system"
	"github.com/TeamPhoneix/go-evil/domains/time"
	"github.com/TeamPhoneix/go-evil/domains/webview"
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
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
		call_function := ""

		// Identify which domain to call on
		regex := regexp.MustCompile(DOMAIN_FUNC_VALUE)
		data := regex.FindAllStringSubmatch(line, -1)

		if len(data) > 0 {
			// This makes it easier to figure out what is what
			domain := data[0][1]
			function := data[0][2]
			value := data[0][3]

			call_function, s_json = grab_code(domain, function, value, s_json)

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
				call_function, s_json = construct_foreach_loop(data[0][1], body, s_json)

			}
		}

		if call_function != "" { // Don't want any empty lines
			calls = append(calls, call_function)
		}
	}

	return calls, s_json
}

//
//
// Grab the contents from each domains parser
//
//
func grab_code(domain string, function string, value string, s_json string) (string, string) {
	call_function := ""

	// Translating variables
	regex := regexp.MustCompile(GET_VAR)
	result := regex.FindAllStringSubmatch(value, -1)

	if len(result) > 0 {
		data_object := structure.Receive(s_json)
		var_call := result[0][1]
		Var_type := result[0][2]
		var_id := result[0][3]

		var_value := data_object.Get_variable_value(Var_type, var_id)
		value = strings.ReplaceAll(value, var_call, var_value)

		s_json = structure.Send(data_object)
	}

	switch domain {
	case "system":
		call_function, s_json = system.Parser(function, value, s_json)

	case "time":
		call_function, s_json = time.Parser(function, value, s_json)

	case "webview":
		call_function, s_json = webview.Parser(function, value, s_json)

	case "self":
		call_function, s_json = self.Parser(function, value, s_json)

	}

	return call_function, s_json
}

//
//
// Construcs the code needed for a "foreach" loop
//
//
func construct_foreach_loop(condition string, body []string, s_json string) (string, string) {
	call := "foreach"

	body_calls, s_json := convert_code(body, s_json) // Converts the code for the foreach body
	data_object := structure.Receive(s_json)

	condition = tools.Erase_delimiter(condition, "\"") // Removes all found "

	final_body := []string{fmt.Sprintf(
		"func %s(values string){", call),
		"values = runtime_var.get(values)",
		"array := tools.Extract_values_array(values)",
		"for _, value := range array{",
		"runtime_var.foreach = value"}

	final_body = append(final_body, body_calls...)
	final_body = append(final_body, "}}")

	data_object.Add_go_function(final_body)

	return fmt.Sprintf("%s(\"%s\")", call, condition), structure.Send(data_object)
}

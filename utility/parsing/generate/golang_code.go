package generate

import (
	"regexp"

	"github.com/TeamPhoneix/go-evil/utility/parsing/foreach"
	if_else "github.com/TeamPhoneix/go-evil/utility/parsing/if_else"

	"github.com/TeamPhoneix/go-evil/utility/parsing/imports"
	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"
)

// Converts evil code to golang code and returns it
func Generate_golang_code(gut []string, s_json string) ([]string, string) {
	calls := []string{}

	for i := 0; i < len(gut); i++ {
		line := gut[i]
		call_functions := []string{}

		// Identify which domain to call on
		regex := regexp.MustCompile(evil_regex.DOMAIN_FUNC_VALUE)
		data := regex.FindAllStringSubmatch(line, -1)

		if len(data) > 0 {
			// This makes it easier to figure out what is what
			domain := data[0][1]
			function := data[0][2]
			value := data[0][3]

			call_functions, s_json = imports.Construct_domain_code(domain, function, value, s_json)

		} else {
			regex = regexp.MustCompile(evil_regex.GET_FOREACH_HEADER)
			foreach_identified := regex.FindAllStringSubmatch(line, -1)
			regex = regexp.MustCompile(evil_regex.GET_IF_HEADER)
			if_identified := regex.FindAllStringSubmatch(line, -1)

			if len(foreach_identified) > 0 { // foreach loop
				body := foreach.Get_foreach_body(&i, gut)
				call_functions, s_json = Construct_foreach_loop(foreach_identified[0][1], body, s_json)

			} else if len(if_identified) > 0 { // if/else statement
				true_body, false_body := if_else.Get_if_else_body(&i, gut)
				call_functions, s_json = Construct_if_else(if_identified[0][1], true_body, false_body, s_json)
			}

		}

		if len(call_functions) > 0 { // Don't want any empty lines
			calls = append(calls, call_functions...)
		}
	}

	return calls, s_json
}

package parsing

import (
	"regexp"

	"github.com/TeamPhoneix/go-evil/domains/system"
	"github.com/TeamPhoneix/go-evil/domains/time"
	"github.com/TeamPhoneix/go-evil/utility/structure"
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

	for _, line := range gut {
		// Identify which domain to call on
		regex := regexp.MustCompile(DOMAIN_FUNC_VALUE)
		call := regex.FindAllStringSubmatch(line, -1)

		if len(call) > 0 {
			// This makes it easier to figure out what is what
			domain := call[0][1]
			function := call[0][2]
			value := call[0][3]
			call_function := ""

			call_function, s_json = grab_code(domain, function, value, s_json)

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

	switch domain {
	case "system":
		call_function, s_json = system.Parser(function, value, s_json)

	case "time":
		call_function, s_json = time.Parser(function, value, s_json)
	}

	return call_function, s_json
}

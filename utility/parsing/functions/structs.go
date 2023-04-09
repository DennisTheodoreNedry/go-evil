package functions

import (
	"regexp"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Parses the data from the target file and generates function structures from it
// The gut contains the evil code that later on will be parsed
func Build_functions_structs(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(evil_regex.FUNC)
	functions := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(functions) > 0 {
		for _, function := range functions {
			index := 3

			return_type := "null"
			f_type := function[1]
			name := function[2]

			if f_type == "c" {
				return_type = function[4]
				index = 5
			}

			gut := function[index : len(function)-1]

			data_object.Add_function(name, f_type, return_type, gut)

		}
	}
	return structure.Send(data_object)
}
package configuration

import (
	"regexp"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"
	"github.com/TeamPhoneix/go-evil/utility/structure/json"
)

// Checks if the user specificed a name for the malware
func check_output(line string, data_object *json.Json_t) {

	if data_object.Binary_name == "" { // Don't override if the user already have provided a name
		regex := regexp.MustCompile(evil_regex.CONFIGURATION_NAME)
		result := regex.FindAllStringSubmatch(line, -1)
		name := "me_not_a_virus"

		if len(result) > 0 {
			name = result[0][1]
		}

		data_object.Set_binary_name(name)
	}

}

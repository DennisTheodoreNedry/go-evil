package configuration

import (
	"regexp"
	"strings"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Checks if the user specificed the output to be obfuscated or not
func check_obfuscate(line string, s_json string) string {
	data_object := structure.Receive(s_json)

	if !data_object.Obfuscate { // No point in checking, if the user already has enabled it through CLI
		regex := regexp.MustCompile(evil_regex.CONFIGURATION_OBFUSCATE)
		result := regex.FindAllStringSubmatch(line, -1)

		if len(result) > 0 {
			option := strings.ToLower(result[0][1])
			if option == "true" {
				data_object.Enable_obfuscate()
			}
		}
	}

	return structure.Send(data_object)
}

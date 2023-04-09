package configuration

import (
	"regexp"
	"strings"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Checks if the user specificed behavior if it's being run through a debugger
func check_debugger_behavior(line string, s_json string) string {
	data_object := structure.Receive(s_json)

	if data_object.Debugger_behavior == "" { // No point in checking, if the user already has enabled it through CLI
		regex := regexp.MustCompile(evil_regex.CONFIGURATION_DEBUGGER_BEHAVIOR)
		result := regex.FindAllStringSubmatch(line, -1)
		behavior := "stop"

		if len(result) > 0 {
			behavior = strings.ToLower(result[0][1])
		}

		data_object.Change_detection_behavior(behavior)
	}

	return structure.Send(data_object)
}
package configuration

import (
	"regexp"
	"runtime"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Checks if the user specificed an os for the malware
func check_os(line string, s_json string) string {
	data_object := structure.Receive(s_json)

	if data_object.Target_os == "" { // Don't override if the user already have provided a os
		regex := regexp.MustCompile(evil_regex.CONFIGURATION_OS)
		result := regex.FindAllStringSubmatch(line, -1)
		os := runtime.GOOS

		if len(result) > 0 {
			os = result[0][1]
		}

		data_object.Set_target_os(os)

	}

	return structure.Send(data_object)
}

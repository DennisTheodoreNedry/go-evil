package configuration

import (
	"regexp"
	"runtime"

	evil_regex "github.com/DennisTheodoreNedry/go-evil/utility/parsing/regex"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Checks if the user specificed an os for the malware
func check_os(line string, data_object *json.Json_t) {

	if data_object.Target_os == "" { // Don't override if the user already have provided a os
		regex := regexp.MustCompile(evil_regex.CONFIGURATION_OS)
		result := regex.FindAllStringSubmatch(line, -1)
		os := runtime.GOOS

		if len(result) > 0 {
			os = result[0][1]
		}

		data_object.Set_target_os(os)

	}

}

package configuration

import (
	"regexp"
	"runtime"

	evil_regex "github.com/DennisTheodoreNedry/go-evil/utility/parsing/regex"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Checks if the user specificed an architecture for the malware
func check_architecture(line string, data_object *json.Json_t) {

	if data_object.Target_arch == "" { // Don't override if the user already have provided a name
		regex := regexp.MustCompile(evil_regex.CONFIGURATION_ARCH)
		result := regex.FindAllStringSubmatch(line, -1)
		arch := runtime.GOARCH

		if len(result) > 0 {
			arch = result[0][1]
		}

		data_object.Set_target_arch(arch)

	}

}

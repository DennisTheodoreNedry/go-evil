package configuration

import (
	"regexp"

	evil_regex "github.com/DennisTheodoreNedry/go-evil/utility/parsing/regex"
	"github.com/DennisTheodoreNedry/go-evil/utility/version"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Checks if the provided evil file can be compiled by this compiler
func check_version(line string) {

	regex := regexp.MustCompile(evil_regex.COMPILER_VERSION)
	result := regex.FindAllStringSubmatch(line, -1)
	current_version := version.EVIL_VERSION_SMALL

	if len(result) > 0 {
		grabbed_version := result[0][1]

		if grabbed_version > current_version {
			notify.Error("The provided evil file is targeting a newer version of go-evil and is therefore not supported", "parsing.check_version()", 1)

		} else if grabbed_version < current_version {
			notify.Warning("The provided evil file is targeting an older version of go-evil and therefore might not be supported by this compiler")
		}
	}

}

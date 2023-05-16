package configuration

import (
	"regexp"

	evil_regex "github.com/s9rA16Bf4/go-evil/utility/parsing/regex"
	"github.com/s9rA16Bf4/go-evil/utility/version"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Checks if the provided evil file can be compiled by this compiler
func check_version(line string) {

	regex := regexp.MustCompile(evil_regex.COMPILER_VERSION)
	result := regex.FindAllStringSubmatch(line, -1)
	current_version := version.EVIL_VERSION

	if len(result) > 0 {
		grabbed_version := result[0][1]

		if grabbed_version > current_version {
			notify.Error("The provided Evil file is targeting a newer version of go-evil and is therefore not supported", "parsing.check_version()")

		} else if grabbed_version < current_version {
			notify.Warning("The provided Evil file is targeting an older version of go-evil and therefore might not be supported by this compiler")
		}
	}

}

package version

import (
	"regexp"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EVIL_VERSION     = "1.2.0"           // high version. medium version. small version
	EXTRACT_COMPILER = "([0-9]\\.[0-9])" // Extracts the high and medium version of the compiler
)

func Print_version() {
	notify.Inform("Major version: " + Get_Compiler_version())
	notify.Inform("Current version: " + EVIL_VERSION)
}

func Get_Compiler_version() string {
	regex := regexp.MustCompile(EXTRACT_COMPILER)
	result := regex.FindAllStringSubmatch(EVIL_VERSION, -1)
	return result[0][1]
}

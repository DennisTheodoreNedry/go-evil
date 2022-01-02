package version

import (
	"regexp"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	VERSION                 = "1.0.0"           // high version. medium version. small version
	EXTRACT_HIGH_AND_MEDIUM = "([0-9]\\.[0-9])" // Extracts the high and medium version of the compiler
)

func Print_version() {
	notify.Inform("Current version of the evil compiler is: " + VERSION)
}

func Get_high_medium_version() string {
	regex := regexp.MustCompile(EXTRACT_HIGH_AND_MEDIUM)
	result := regex.FindAllStringSubmatch(VERSION, -1)
	return result[0][1]
}

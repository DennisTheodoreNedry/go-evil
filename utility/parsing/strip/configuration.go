package strip

import (
	"regexp"
	"strings"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Removes the configuration section if it is found
func remove_configuration(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(evil_regex.COMPILER_CONFIGURATION)
	comments := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(comments) > 0 {
		for _, line := range comments[0] {
			data_object.File_gut = strings.Replace(data_object.File_gut, line, "", -1)
		}
	}

	return structure.Send(data_object)
}

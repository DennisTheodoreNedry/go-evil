package strip

import (
	"regexp"
	"strings"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Finds all comments and removes them
func remove_injected_headers(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(evil_regex.INJECTION_GO_HEADERS)
	comments := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(comments) > 0 {
		for _, line := range comments[0] {
			data_object.File_gut = strings.Replace(data_object.File_gut, line, "", -1)
		}
	}

	return structure.Send(data_object)
}

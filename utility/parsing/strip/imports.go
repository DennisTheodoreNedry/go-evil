package strip

import (
	"fmt"
	"regexp"
	"strings"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Removes all imports from the structure
func remove_imports(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(evil_regex.IMPORT)
	comments := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(comments) > 0 {
		for _, line := range comments[0] {
			data_object.File_gut = strings.Replace(data_object.File_gut, fmt.Sprintf("use %s", line), "", -1)
		}
	}

	return structure.Send(data_object)
}

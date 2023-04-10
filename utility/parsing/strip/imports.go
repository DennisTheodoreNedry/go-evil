package strip

import (
	"fmt"
	"regexp"
	"strings"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"
	"github.com/TeamPhoneix/go-evil/utility/structure/json"
)

// Removes all imports from the structure
func remove_imports(data_object *json.Json_t) {
	regex := regexp.MustCompile(evil_regex.IMPORT)
	comments := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(comments) > 0 {
		for _, line := range comments[0] {
			data_object.File_gut = strings.Replace(data_object.File_gut, fmt.Sprintf("use %s", line), "", -1)
		}
	}

}

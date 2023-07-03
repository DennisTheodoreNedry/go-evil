package evilarray

import (
	"regexp"
	"strings"

	"github.com/s9rA16Bf4/Go-tools/tools"
)

// Parses the provided evil array and inserts those values into this structure
func (object *Evil_array_t) Parse(formatted_evil_array string) {

	regex := regexp.MustCompile(EXTRACT_VALUES_FROM_EVIL_ARRAY)
	values := regex.FindAllStringSubmatch(formatted_evil_array, -1)

	if len(values) > 0 {
		for _, line := range strings.Split(values[0][1], ",") {
			result := tools.Starts_with(line, []string{" "})

			if ok := result[" "]; ok { // It begins with a space
				line = tools.Erase_delimiter(line, []string{" "}, 1)
			}

			line = tools.Erase_delimiter(line, []string{"\""}, -1)
			object.gut = append(object.gut, line)
			object.length++
		}
	}
}

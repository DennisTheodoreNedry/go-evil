package strip

import (
	"regexp"
	"strings"

	evil_regex "github.com/DennisTheodoreNedry/go-evil/utility/parsing/regex"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Finds all comments and removes them
func remove_injected_code(data_object *json.Json_t) {

	regex := regexp.MustCompile(evil_regex.INJECTION_GO_CODE)
	comments := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(comments) > 0 {
		for _, line := range comments[0] {
			data_object.File_gut = strings.Replace(data_object.File_gut, line, "", -1)
		}
	}

}

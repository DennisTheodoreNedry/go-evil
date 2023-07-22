package imports

import (
	"regexp"

	evil_regex "github.com/s9rA16Bf4/go-evil/utility/parsing/regex"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Adds all found imports
func Find_imports(data_object *json.Json_t) {
	regex := regexp.MustCompile(evil_regex.IMPORT)
	result := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(result) > 0 {
		for _, domain := range result {
			data_object.Add_domain(domain[1])
		}
	}

}

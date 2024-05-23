package error

import (
	"fmt"
	"regexp"

	evil_regex "github.com/DennisTheodoreNedry/go-evil/utility/parsing/regex"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Checks that a used domain has been imported
func check_imports(data_object *json.Json_t) {

	regex := regexp.MustCompile(evil_regex.IMPORT)
	domains := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	regex = regexp.MustCompile(evil_regex.DOMAIN_FUNC_VALUE)
	calls := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	for _, call := range calls {
		found_domain := false
		for _, domain := range domains {
			if call[1] == domain[1] {
				found_domain = true
				break
			}
		}

		if !found_domain {
			notify.Error(fmt.Sprintf("The domain '%s' was used but were never imported!", call[1]), "error.check_imports()", 1)
		}
	}
}

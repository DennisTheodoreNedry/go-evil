package error

import (
	"fmt"
	"regexp"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Checks that a used domain has been imported
func check_imports(s_json string) {
	data_object := structure.Receive(s_json)
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
			notify.Error(fmt.Sprintf("The domain '%s' was used but were never imported!", call[1]), "error.check_imports()")
		}
	}
}

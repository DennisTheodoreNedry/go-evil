package parsing

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Check_for_errors(s_json string) {

	comments(s_json)
	detect_functions(s_json)
	check_imports(s_json)
}

//
//
// Checks for comments that have not been terminated
//
//
func comments(s_json string) {
	data_object := structure.Receive(s_json)

	gut := strings.Split(data_object.File_gut, "\n")
	regex_pattern := []string{COMMENT_WRONG_LHS, COMMENT_WRONG_RHS}

	for i, line := range gut {
		line = strings.Replace(line, " ", "", -1)  // Remove every space
		line = strings.Replace(line, "\t", "", -1) // and remove every tab, this to ensure that we our regex is gonna work correctly

		for _, pattern := range regex_pattern {
			regex := regexp.MustCompile(pattern)
			result := regex.FindAllStringSubmatch(line, -1)

			if len(result) != 0 {
				notify.Error(fmt.Sprintf("Found a wrongly formatted comment on line %d,", i+1), "error.comments()")
			}
		}

	}
}

//
//
// Detects if there are any functions in the file
//
//
func detect_functions(s_json string) {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(FUNC)
	functions := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(functions) == 0 { // No functions were detected
		notify.Error(fmt.Sprintf("No functions were found in the file '%s'", data_object.File_path), "error.detect_functions()")
	}
}

//
//
// Checks that a used domain has been imported
//
//
func check_imports(s_json string) {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(IMPORT)
	domains := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	regex = regexp.MustCompile(DOMAIN_FUNC_VALUE)
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

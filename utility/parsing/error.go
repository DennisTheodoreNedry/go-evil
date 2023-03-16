package parsing

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Wrapper function which calls all our error checking functions
//
//
func Check_for_errors(s_json string) {

	comments(s_json)
	detect_functions(s_json)
	check_imports(s_json)
	check_strings(s_json)
	check_evil_arrays(s_json)
	check_compile_variable(s_json)
	check_runtime_variable(s_json)
	check_call_function_format(s_json)
}

//
//
// Checks for comments that have not been terminated
//
//
func comments(s_json string) {
	data_object := structure.Receive(s_json)

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		result := strings.Count(line, "@")

		if result%2 != 0 {
			notify.Error(fmt.Sprintf("Found a wrongly formatted string on line %d\nError line: '%s'", i+1, line), "error.check_strings()")
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

//
//
// Checks for strings that have not been terminated
//
//
func check_strings(s_json string) {
	data_object := structure.Receive(s_json)

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		// We need to check so that the line doesn't start with a comment
		comment_status := tools.Starts_with(line, []string{"@"})
		ok := comment_status["@"]

		bunny_ears := strings.Count(line, "\"")
		if bunny_ears%2 != 0 && !ok {
			notify.Error(fmt.Sprintf("Found a wrongly formatted string on line %d\nError line: '%s'", i+1, line), "error.check_strings()")
		}

	}
}

//
//
// Checks for arrays that have not been terminated
//
//
func check_evil_arrays(s_json string) {
	data_object := structure.Receive(s_json)

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		l_wing := strings.Count(line, "${")
		r_wing := strings.Count(line, "}$")

		if l_wing != r_wing {
			notify.Error(fmt.Sprintf("Found a wrongly formatted string on line %d\nError line: '%s'", i+1, line), "error.check_evil_arrays()")
		}

	}
}

//
//
// Checks for compile variables that have not been terminated
//
//
func check_compile_variable(s_json string) {
	data_object := structure.Receive(s_json)

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		count := strings.Count(line, "$")

		if count%2 != 0 {
			notify.Error(fmt.Sprintf("Found a wrongly formatted string on line %d\nError line: '%s'", i+1, line), "error.check_evil_arrays()")
		}

	}
}

//
//
// Checks for runtime variables that have not been terminated
//
//
func check_runtime_variable(s_json string) {
	data_object := structure.Receive(s_json)

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		count := strings.Count(line, "€")

		if count%2 != 0 {
			notify.Error(fmt.Sprintf("Found a wrongly formatted string on line %d\nError line: '%s'", i+1, line), "error.check_evil_arrays()")
		}

	}
}

//
//
// Checks if potential call functions are wrongly formatted
//
//
func check_call_function_format(s_json string) {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(FUNC)
	functions := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	for _, d_func := range functions {
		if d_func[1] == "c" && !strings.Contains(d_func[3], "->") {
			notify.Error(fmt.Sprintf("Found a wrongly formatted function with the name '%s'", d_func[2]), "error.check_call_function_format()")
		}
	}
}

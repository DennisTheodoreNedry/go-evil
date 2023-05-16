package error

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/parsing/regex"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Detects if there are any functions in the file
func detect_functions(data_object *json.Json_t) {

	regex := regexp.MustCompile(regex.FUNC)
	functions := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(functions) == 0 { // No functions were detected
		notify.Error(fmt.Sprintf("No functions were found in the file '%s'", data_object.File_path), "error.detect_functions()")
	}
}

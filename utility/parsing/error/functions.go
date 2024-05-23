package error

import (
	"fmt"
	"regexp"

	"github.com/DennisTheodoreNedry/go-evil/utility/parsing/regex"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Detects if there are any functions in the file
func detect_functions(data_object *json.Json_t) {

	regex := regexp.MustCompile(regex.FUNC)
	functions := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(functions) == 0 { // No functions were detected
		notify.Error(fmt.Sprintf("No functions were found in the file '%s'", data_object.File_path), "error.detect_functions()", 1)
	}
}

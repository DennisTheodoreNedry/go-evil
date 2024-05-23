package error

import (
	"fmt"
	"strings"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Checks for strings that have not been terminated
func check_strings(data_object *json.Json_t) {

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		// We need to check so that the line doesn't start with a comment
		comment_status := gotools.StartsWith(line, []string{"@"})
		ok := comment_status["@"]

		bunny_ears := strings.Count(line, "\"")
		if bunny_ears%2 != 0 && !ok {
			notify.Error(fmt.Sprintf("Found a wrongly formatted string on line %d\nError line: '%s'", i+1, line), "error.check_strings()", 1)
		}

	}
}

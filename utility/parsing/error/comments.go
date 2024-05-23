package error

import (
	"fmt"
	"strings"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Checks for comments that have not been terminated
func check_comments(data_object *json.Json_t) {

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		result := strings.Count(line, "@")

		if result%2 != 0 {
			notify.Error(fmt.Sprintf("Found a wrongly formatted comment on line %d\nError line: '%s'", i+1, line), "error.check_comments()", 1)
		}

	}
}

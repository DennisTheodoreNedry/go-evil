package error

import (
	"fmt"
	"strings"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Checks for arrays that have not been terminated
func check_evil_arrays(data_object *json.Json_t) {

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		l_wing := strings.Count(line, "${")
		r_wing := strings.Count(line, "}$")

		if l_wing != r_wing {
			notify.Error(fmt.Sprintf("Found a wrongly formatted evil array on line %d\nError line: '%s'", i+1, line), "error.check_evil_arrays()", 1)
		}

	}
}

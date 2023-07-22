package error

import (
	"fmt"
	"strings"

	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
)

// Checks for runtime variables that have not been terminated
func check_runtime_variable(data_object *json.Json_t) {

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		count := strings.Count(line, "€")

		if count%2 != 0 {
			notify.Error(fmt.Sprintf("Found a wrongly formatted runtime variable on line %d\nError line: '%s'", i+1, line), "error.check_runtime_variable()", 1)
		}

	}
}

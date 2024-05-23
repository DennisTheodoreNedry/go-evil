package error

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/DennisTheodoreNedry/go-evil/utility/parsing/regex"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Checks if potential call functions are wrongly formatted
func check_call_function_format(data_object *json.Json_t) {

	regex := regexp.MustCompile(regex.FUNC)
	functions := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	for _, d_func := range functions {
		if d_func[1] == "c" && !strings.Contains(d_func[3], "->") {
			notify.Error(fmt.Sprintf("Found a wrongly formatted function with the name '%s'", d_func[2]), "error.check_call_function_format()", 1)
		}
	}
}

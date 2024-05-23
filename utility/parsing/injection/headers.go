package injection

import (
	"regexp"
	"strings"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	evil_regex "github.com/DennisTheodoreNedry/go-evil/utility/parsing/regex"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

func Grab_injected_headers(data_object *json.Json_t) {

	regex := regexp.MustCompile(evil_regex.INJECTION_GO_HEADERS)
	result := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(result) == 0 {
		data_object.Log_object.Log("Found no injected headers", 2)

	} else {
		headers := strings.Split(result[0][1], "\n")

		for _, header := range headers {
			header = gotools.EraseDelimiter(header, []string{" "}, -1) // Erase any potential spaces
			data_object.Add_go_import(header)
		}
	}

}

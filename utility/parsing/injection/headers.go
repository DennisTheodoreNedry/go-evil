package injection

import (
	"regexp"
	"strings"

	evil_regex "github.com/s9rA16Bf4/go-evil/utility/parsing/regex"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Grab_injected_headers(data_object *json.Json_t) {

	regex := regexp.MustCompile(evil_regex.INJECTION_GO_HEADERS)
	result := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(result) == 0 {
		notify.Log("Found no injected headers", data_object.Verbose_lvl, "2")

	} else {
		headers := strings.Split(result[0][1], "\n")

		for _, header := range headers {
			header = tools.Erase_delimiter(header, []string{" "}, -1) // Erase any potential spaces
			data_object.Add_go_import(header)
		}
	}

}

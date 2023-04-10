package injection

import (
	"regexp"
	"strings"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Grab_injected_headers(s_json string) string {
	data_object := structure.Receive(s_json)

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

	return structure.Send(data_object)
}

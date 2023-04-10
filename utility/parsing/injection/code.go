package injection

import (
	"regexp"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Grab_injected_code(s_json string) string {
	data_object := structure.Receive(s_json)

	regex := regexp.MustCompile(evil_regex.INJECTION_GO_CODE)
	result := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(result) == 0 {
		notify.Log("Found no code injection", data_object.Verbose_lvl, "2")

	} else {
		for _, injected_function := range result {
			func_type := injected_function[1]
			func_name := tools.Generate_random_n_string(8)
			func_gut := injected_function[2 : len(injected_function)-1]

			data_object.Add_function(func_name, func_type, "null", func_gut)

		}
	}

	return structure.Send(data_object)
}

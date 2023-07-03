package injection

import (
	"regexp"

	"github.com/s9rA16Bf4/Go-tools/tools"
	evil_regex "github.com/s9rA16Bf4/go-evil/utility/parsing/regex"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Grab_injected_code(data_object *json.Json_t) {

	regex := regexp.MustCompile(evil_regex.INJECTION_GO_CODE)
	result := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(result) == 0 {
		notify.Log("Found no code injection", data_object.Verbose_lvl, "2")

	} else {
		for _, injected_function := range result {
			func_type := injected_function[1]
			func_name := tools.Generate_random_n_string(8)
			func_gut := injected_function[2 : len(injected_function)-1]

			// Let's identify which type of function type this is
			switch func_type {
			case "l":
				data_object.Add_loop_function(func_name)

			case "b":
				data_object.Add_boot_function(func_name)

			case "e":
				data_object.Add_end_function(func_name)
			}

			data_object.Add_go_function(functions.Go_func_t{Name: func_name, Func_type: "", Part_of_struct: "", Return_type: "", Parameters: []string{}, Gut: func_gut})

		}
	}

}

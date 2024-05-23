package injection

import (
	"regexp"
	"strings"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	evil_regex "github.com/DennisTheodoreNedry/go-evil/utility/parsing/regex"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

func Grab_injected_code(data_object *json.Json_t) {

	regex := regexp.MustCompile(evil_regex.INJECTION_GO_CODE)
	result := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(result) == 0 {
		data_object.Log_object.Log("Found no code injection", 2)

	} else {
		for _, injected_function := range result {
			func_type := injected_function[1]
			func_name := gotools.Generate_random_n_string(8)
			func_gut := injected_function[2 : len(injected_function)-1]

			// Let's identify which type of function type this is
			switch strings.ToLower(func_type) {
			case "loop":
				data_object.Add_loop_function(func_name)

			case "boot":
				data_object.Add_boot_function(func_name)

			case "end":
				data_object.Add_end_function(func_name)
			}

			data_object.Add_go_function(functions.Go_func_t{Name: func_name, Func_type: "", Part_of_struct: "", Return_type: "", Parameters: []string{}, Gut: func_gut})

		}
	}

}

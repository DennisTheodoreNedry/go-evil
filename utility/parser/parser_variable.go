package parser

import (
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/json"
	compiler_time "github.com/s9rA16Bf4/go-evil/utility/variables/compiler-time"
)

const (
	EXTRACT_COMP_VARIABLE = "(\\$[0-9]+)" // Extracts a compile time variable
	EXTRACT_RUN_VARIABLE  = "(€[0-9]+)"   // Extracts a run time variable
)

func Variable(base_64_serialize_json string) string {

	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("Parser.Variable()")

	var new_line []string
	regex_comp := regexp.MustCompile(EXTRACT_COMP_VARIABLE)
	regex_run := regexp.MustCompile(EXTRACT_RUN_VARIABLE)

	for _, line := range data_structure.File_gut {
		result := regex_comp.FindAllStringSubmatch(line, -1)
		if len(result) > 0 {
			data_structure.Append_compile_time_var(result[0][1])
			data_structure.Append_compile_time_value(compiler_time.Get_variable(result[0][1]))
		}

		result = regex_run.FindAllStringSubmatch(line, -1)
		if len(result) > 0 {
			data_structure.Append_run_time_var(result[0][1])
		}

		new_line = append(new_line, compiler_time.Get_variable(line))
	}

	data_structure.File_gut = new_line

	return json.Send(data_structure)
}

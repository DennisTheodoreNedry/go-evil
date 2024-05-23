package error

import (
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Wrapper function which calls all our error checking functions
func Check_for_errors(data_object *json.Json_t) {

	check_comments(data_object)
	detect_functions(data_object)
	check_imports(data_object)
	check_strings(data_object)
	check_evil_arrays(data_object)
	check_compile_variable(data_object)
	check_runtime_variable(data_object)
	check_call_function_format(data_object)
}

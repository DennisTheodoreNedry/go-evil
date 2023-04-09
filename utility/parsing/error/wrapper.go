package error

// Wrapper function which calls all our error checking functions
func Check_for_errors(s_json string) {

	comments(s_json)
	detect_functions(s_json)
	check_imports(s_json)
	check_strings(s_json)
	check_evil_arrays(s_json)
	check_compile_variable(s_json)
	check_runtime_variable(s_json)
	check_call_function_format(s_json)
}

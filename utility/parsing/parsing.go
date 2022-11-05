package parsing

//
//
// preface before we start parsing
//
//
func preface(s_json string) string {
	Check_for_errors(s_json)                 // Checks for common errors found in the file
	s_json = Check_configuration(s_json)     // Checks for a configuration setting in the file
	s_json = Find_imports(s_json)            // Finds all imports in the file
	s_json = Strip(s_json)                   // Removes the configuration section and every comment found
	s_json = Build_functions_structs(s_json) // Builds function structs for each found function
	return s_json
}

//
//
// Parses the contents of the provided file
//
//
func Parse(s_json string) string {
	s_json = preface(s_json) // Handles every preface we could possibly want done before we start parsing

	s_json = generate_structs(s_json)
	s_json = generate_behavior_debugging(s_json)

	s_json, boot_func, loop_func := generate_go_functions(s_json)

	s_json = generate_main_function(s_json, boot_func, loop_func)

	s_json = construct_final_malware(s_json)

	return s_json
}

package parsing

import (
	evil_configuration "github.com/TeamPhoneix/go-evil/utility/parsing/configuration"
	evil_error "github.com/TeamPhoneix/go-evil/utility/parsing/error"
	evil_functions "github.com/TeamPhoneix/go-evil/utility/parsing/functions"
	evil_import "github.com/TeamPhoneix/go-evil/utility/parsing/imports"
	evil_inject "github.com/TeamPhoneix/go-evil/utility/parsing/injection"
	evil_strip "github.com/TeamPhoneix/go-evil/utility/parsing/strip"
)

// preface before we start parsing
func preface(s_json string) string {
	evil_error.Check_for_errors(s_json)                     // Checks for common errors found in the file
	s_json = evil_configuration.Check_configuration(s_json) // Checks for a configuration setting in the file
	s_json = evil_import.Find_imports(s_json)               // Finds all imports in the file
	s_json = evil_inject.Grab_injected_headers(s_json)      // Finds potential headers and removes the section if it's found
	s_json = evil_inject.Grab_injected_code(s_json)         // Finds potential code and removes the section if it's found
	s_json = evil_strip.Strip(s_json)                       // Removes the configuration section and every comment found

	s_json = evil_functions.Build_functions_structs(s_json) // Builds function structs for each found function

	return s_json
}

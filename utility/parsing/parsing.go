package parsing

import (
	evil_configuration "github.com/TeamPhoneix/go-evil/utility/parsing/configuration"
	"github.com/TeamPhoneix/go-evil/utility/parsing/debugger"
	evil_error "github.com/TeamPhoneix/go-evil/utility/parsing/error"
	evil_final "github.com/TeamPhoneix/go-evil/utility/parsing/finalize"
	evil_functions "github.com/TeamPhoneix/go-evil/utility/parsing/functions"
	evil_generate "github.com/TeamPhoneix/go-evil/utility/parsing/generate"
	evil_import "github.com/TeamPhoneix/go-evil/utility/parsing/imports"
	evil_strip "github.com/TeamPhoneix/go-evil/utility/parsing/strip"
)

// preface before we start parsing
func preface(s_json string) string {
	evil_error.Check_for_errors(s_json)                     // Checks for common errors found in the file
	s_json = evil_configuration.Check_configuration(s_json) // Checks for a configuration setting in the file
	s_json = evil_import.Find_imports(s_json)               // Finds all imports in the file
	s_json = evil_strip.Strip(s_json)                       // Removes the configuration section and every comment found

	s_json = evil_functions.Build_functions_structs(s_json) // Builds function structs for each found function

	return s_json
}

// Parses the contents of the provided file
func Parse(s_json string) string {
	s_json = preface(s_json) // Handles every preface we could possibly want done before we start parsing

	s_json = evil_generate.Generate_structs(s_json)

	s_json = debugger.Generate_behavior(s_json)

	s_json, boot_func, loop_func := evil_generate.Generate_go_functions(s_json)

	s_json = evil_generate.Generate_main(s_json, boot_func, loop_func)

	s_json = evil_final.Construct_malware(s_json)

	return s_json
}

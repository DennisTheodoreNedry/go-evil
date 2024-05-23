package parsing

import (
	"github.com/DennisTheodoreNedry/go-evil/utility/io"
	evil_configuration "github.com/DennisTheodoreNedry/go-evil/utility/parsing/configuration"
	evil_error "github.com/DennisTheodoreNedry/go-evil/utility/parsing/error"
	evil_functions "github.com/DennisTheodoreNedry/go-evil/utility/parsing/functions"
	evil_import "github.com/DennisTheodoreNedry/go-evil/utility/parsing/imports"
	evil_inject "github.com/DennisTheodoreNedry/go-evil/utility/parsing/injection"
	evil_strip "github.com/DennisTheodoreNedry/go-evil/utility/parsing/strip"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// preface before we start parsing
func preface(data_object *json.Json_t) {
	io.Read_file(data_object) // Read the file

	evil_error.Check_for_errors(data_object)            // Checks for common errors found in the file
	evil_configuration.Check_configuration(data_object) // Checks for a configuration setting in the file
	evil_import.Find_imports(data_object)               // Finds all imports in the file
	evil_inject.Grab_injected_headers(data_object)      // Finds potential headers and removes the section if it's found
	evil_inject.Grab_injected_code(data_object)         // Finds potential code and removes the section if it's found
	evil_strip.Strip(data_object)                       // Removes the configuration section and every comment found

	evil_functions.Build_functions_structs(data_object) // Builds function structs for each found function

}

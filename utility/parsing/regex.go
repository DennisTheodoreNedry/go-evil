package parsing

//
//
// This file contains all regex patterns utilized
//
//

const (
	FUNC              = "([blc]) ([a-z0-9_]+) *{\n* *((.*|\n*)*?)\n*}" // Extracts all functions
	DOMAIN_FUNC_VALUE = "([a-z_]+)::([a-z_]+)\\((.*)\\)"               // Extracts the domain, function being called and if a value was sent with it

	IMPORT = "use ([a-z]+)" // Finds imports

	COMMENT = "@.+@" // Identifies a comment

	// Variables //
	GET_VAR      = "(([\\$])([0-9]+)[\\$])" // Grabs the variable type and id
	GET_USER_VAR = "\\$666\\$"              // Grabs what the user is called

	// Configurations //
	COMPILER_CONFIGURATION  = "\\[\n*(?s)(.*)\n*\\]"       // Grabs the configuration secton
	COMPILER_VERSION        = "version +([0-9]+\\.[0-9]+)" // Grabs the compiler that the scrip was meant for
	CONFIGURATION_NAME      = "output +(.*)"
	CONFIGURATION_ARCH      = "arch +(.*)"
	CONFIGURATION_OS        = "os +(.*)"
	CONFIGURATION_EXTENSION = "extension +(.*)"
	CONFIGURATION_OBFUSCATE = "obfuscate +(.*)"
)

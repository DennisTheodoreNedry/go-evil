package regex

//
//
// This file contains all regex patterns utilized
//
//

const (
	// Functions //
	FUNC              = "(boot|loop|call|end) ([a-z0-9_]+) *(-> (string|boolean|integer|null|nil|none))? *{\n* *((.*|\n*)*?)\n*}" // Extracts all functions
	DOMAIN_FUNC_VALUE = "([a-z_#0-9]+)::([a-z_0-9]+)\\(((?s:.)*)\\)"                                                              // Extracts the domain, function being called and if a value was sent with it

	// Imports //
	IMPORT = "use (.*)" // Finds imports

	// Comments //
	COMMENT = "@.+@" // Identifies a comment

	// Variables //
	GET_VAR      = "(\\$([0-9]+)\\$)" // Grabs the variable type and id
	GET_USER_VAR = "\\$666\\$"        // Grabs what the user is called

	// Foreach loops //
	GET_FOREACH_HEADER = "foreach *\\((.+)\\) *:"
	GET_FOREACH_FOOTER = "end foreach"

	// If/else statements //
	GET_IF_HEADER      = "if *\\((.+)\\) *:"
	GET_ELSE_HEADER    = "else:"
	GET_IF_ELSE_FOOTER = "end if"

	// Configurations //
	COMPILER_CONFIGURATION          = "\\[\n*(?s)(.*)\n*\\]"       // Grabs the configuration secton
	COMPILER_VERSION                = "version +([0-9]+\\.[0-9]+)" // Grabs the compiler that the scrip was meant for
	CONFIGURATION_NAME              = "output +(.*)"               // The output name of the binary file
	CONFIGURATION_ARCH              = "arch +(.*)"                 // Target architecture
	CONFIGURATION_OS                = "os +(.*)"                   // Target operating system
	CONFIGURATION_EXTENSION         = "extension +(.*)"            // Extension of the binary file
	CONFIGURATION_OBFUSCATE         = "obfuscate +(.*)"            // Obfuscating the binary file
	CONFIGURATION_DEBUGGER_BEHAVIOR = "debugger_behavior +(.*)"    // Decides how the malware will act if it detects a debugger

	// Injections //
	INJECTION_GO_CODE    = "% *(boot|loop|call|end) *{\n* *((.*|\n*)*?)\n*} *%" // Grab golang code to inject into your application
	INJECTION_GO_HEADERS = "% *\\[\n* *((.*|\n*)*?)\n*\\] *%"                   // Grab potential golang imports
)

package parser

import (
	"regexp"
	"strings"

	attack_vector "github.com/s9rA16Bf4/go-evil/domains/attack_vector"
	"github.com/s9rA16Bf4/go-evil/domains/backdoor"
	"github.com/s9rA16Bf4/go-evil/domains/keyboard"
	"github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/domains/network"
	"github.com/s9rA16Bf4/go-evil/domains/system"
	"github.com/s9rA16Bf4/go-evil/domains/window"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/version"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_MAIN_FUNC           = "((main ?: ?{{1,1}(?s).*}))"                                    // Grabs the main function
	EXTRACT_MAIN_FUNC_HEADER    = "(main:{)"                                                      // We use this to identify if there are multiple main functions in the same file
	EXTRACT_FUNCTION_CALL       = "([a-z]+).*\\((.*)\\);"                                         // Grabs function and a potential value
	EXTRACT_FUNCTION_CALL_WRONG = "([@#a-z]+)\\.([$a-z0-9_]+).([$a-z0-9_]+)\\((\"(.*)\")?\\)[^;]" // And this is utilized to find rows that don't end in ;
	EXTRACT_COMPILER_VERSION    = "\\[.?version +([0-9]+\\.[0-9]+).?\\]"                          // Extracts the major version
	EXTRACT_SYSTEM_VARIABLE     = "(\\$[0-9]+)"                                                   // Extracts the system variable
)

func Interpeter(file_to_read string) {
	content := io.Read_file(file_to_read)

	regex := regexp.MustCompile(EXTRACT_MAIN_FUNC)
	main_function := regex.FindAllStringSubmatch(content, -1)

	if len(main_function) == 0 { // No main function was found
		notify.Error("Failed to find a main function in the provided file "+file_to_read, "parser.interpeter()")
	}

	regex = regexp.MustCompile(EXTRACT_COMPILER_VERSION) // Extracts the high and medium version
	compiler_version := regex.FindAllStringSubmatch(content, -1)
	if len(compiler_version) == 0 { // Compiler version was never specified
		notify.Error("No compiler version was specificed", "parser.interpeter()")
	} else {
		listed_version := compiler_version[0][1]
		if version.Get_Compiler_version() < listed_version {
			notify.Error("Unknown compiler version "+listed_version, "parser.interpeter()")
		} else if version.Get_Compiler_version() > listed_version {
			notify.Warning("You're running a script for an older version of the compiler.\nThis means that there might be functions/syntaxes that have changed")
		}
	}

	regex = regexp.MustCompile(EXTRACT_MAIN_FUNC_HEADER)
	main_header := regex.FindAllStringSubmatch(content, -1)
	if len(main_header) > 1 { // Multiple main functions were defined
		notify.Error("Found multiple main definitions in the provided file "+file_to_read, "parser.interpeter()")
	}
	regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL_WRONG)
	match := regex.FindAllStringSubmatch(content, -1)
	if len(match) > 0 {
		line := match[0][0]
		line = strings.ReplaceAll(line, "\n", "")
		notify.Error("The line '"+line+"' in the file "+file_to_read+" is missing a semi-colon", "parser.interpeter()")
	}

	regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL)
	match = regex.FindAllStringSubmatch(content, -1)
	for _, funct := range match {
		notify.Log("Found possible domain "+funct[1], notify.Verbose_lvl, "3")
		if funct[1][0] == '@' { // Found a comment at the start, so we will ignore this one
			notify.Log("Found comment, will ignore this line", notify.Verbose_lvl, "3")
			continue
		}

		//regex = regexp.MustCompile(EXTRACT_SYSTEM_VARIABLE)
		//variable := regex.FindAllStringSubmatch(funct[4], -1)
		//if len(variable) > 0 { // We found a variable
		//	funct[4] = strings.Replace(funct[4], variable[0][1], system_variable.Get_variable(variable[0][1]), 1) // so we replace it with it's value
		//	notify.Log("Found variable "+variable[0][1]+" which contained the value "+system_variable.Get_variable(variable[0][1]), notify.Verbose_lvl, "2")
		//}
		switch funct[1] { // This will be the top level domain
		case "window":
			io.Append_domain("window")
			window.Parse(funct[0])
		case "system", "#wait":
			io.Append_domain("system")
			system.Parse(funct[0])
		case "network", "#net":
			io.Append_domain("network")
			network.Parse(funct[0])
		case "malware", "#object", "#self", "#this":
			io.Append_domain("malware")
			malware.Parse(funct[0])
		case "keyboard":
			io.Append_domain("keyboard")
			keyboard.Parse(funct[0])
		case "backdoor":
			io.Append_domain("backdoor")
			backdoor.Parse(funct[0])
		case "attack":
			attack_vector.Parse(funct[0])
		}
	}
}

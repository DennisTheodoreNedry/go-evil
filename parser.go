package main

import (
	"regexp"
	"strconv"
)

const EXTRACT_MAIN_FUNC = "main ?: ?{{1,1}(?s).*}{1,1}"                               // Grabs the main function
const EXTRACT_FUNCTION_CALL = "([a-z]+)\\.([a-z]+)\\((\"([A-Za-z0-9 !:.,/]+)\")?\\);" // Grabs function and a potential value
const EXTRACT_EXIT = "system\\.exit\\(\"[0-9]+\"\\);"                                 // We check so that the script actually has an exit statement

func interpeter(file_to_read string) {
	content := read_file(file_to_read)

	regex := regexp.MustCompile(EXTRACT_MAIN_FUNC)
	main_function := regex.FindAllStringSubmatch(content, -1)
	if len(main_function) == 0 { // No main function was found
		notify_error("Failed to find a main function in the provided file "+file_to_read, "parser.interpeter()")
	}
	regex = regexp.MustCompile(EXTRACT_EXIT)
	match := regex.FindAllStringSubmatch(content, -1)
	if len(match) == 0 {
		notify_error("Failed to find an exit statement in the provided file "+file_to_read, "parser.interpeter()")
	}

	// Continue with extracing necessary data
	for i, main_func := range main_function {
		if i > 0 { // Multiple main functions defined (DOESN'T WORk)
			notify_error("Found multiple main definitions in the provided file "+file_to_read, "parser.interpeter()")
		}

		// Extract functions
		for _, line := range main_func {
			regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL)
			match = regex.FindAllStringSubmatch(line, -1)
			for _, funct := range match {
				switch funct[1] {

				case "window": // The window domain was called
					save_domains("window.go")
					switch funct[2] { // Checks the function that were called from the domain
					case "x", "y":
						_, err := strconv.Atoi(funct[4])
						if err != nil {
							notify_error("Failed to convert "+funct[4]+" to integer", "parser.interpreter()")
						}
						if funct[2] == "x" {
							//window_setX(value)
							malw.content = append(malw.content, "window_setX("+funct[4]+")")
						} else {
							//window_setY(value)
							malw.content = append(malw.content, "window_setY("+funct[4]+")")
						}
					case "title":
						//window_setTitle(funct[4])
						malw.content = append(malw.content, "window_setTitle(\""+funct[4]+"\")")

					case "url":
						//window_setDst(funct[4])
						malw.content = append(malw.content, "window_setDst(\""+funct[4]+"\")")

					case "run":
						//window_run()
						malw.content = append(malw.content, "window_run()")

					default:
						notify_error("Unknown function "+funct[2]+" in domain "+funct[1], "parser.interpreter()")
					}

				case "system": // The system domain was called
					save_domains("system.go")
					switch funct[2] { // Function within this domain
					case "exit":
						_, err := strconv.Atoi(funct[4])
						if err != nil {
							notify_error("Failed to convert "+funct[4]+" to integer", "parser.interpreter()")
						}
						malw.content = append(malw.content, "system_exit("+funct[4]+")")

					default:
						notify_error("Unknown function "+funct[2]+" in domain "+funct[1], "parser.interpreter()")
					}

				case "malware":
					save_domains("malware.go")
					switch funct[2] {
					case "name":
						malware_setBinaryName(funct[4])

					default:
						notify_error("Unknown function "+funct[2]+" in domain "+funct[1], "parser.interpreter()")
					}

				default:
					notify_error("Unknown domain "+funct[2], "parser.interpeter()")
				}
			}
		}
	}
}

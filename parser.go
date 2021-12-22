package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const EXTRACT_MAIN_FUNC = "main ?: ?{{1,1}(?s).*}{1,1}"
const EXTRACT_FUNCTION_CALL = "([a-z]+)\\.([a-z]+)\\((\"([A-Za-z0-9 !:.,/]+)\")?\\);" // Grabs function and a potential value

func interpeter(file_to_read string) {
	content := read_file(file_to_read)

	regex := regexp.MustCompile(EXTRACT_MAIN_FUNC)
	match := regex.FindAllStringSubmatch(content, -1)
	if len(match) == 0 { // No main function was found
		notify_error("Failed to find a main function in the provided file <file_name>", "parser.interpeter()")
	}

	// Continue with extracing necessary data
	for i, main_func := range match {
		if i > 0 { // Multiple main functions defined (DOESN'T WORk)
			notify_error("Found multiple main definitions in the provided file <file_name>", "parser.interpeter()")
		}

		// Extract functions
		for _, line := range main_func {
			regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL)
			match = regex.FindAllStringSubmatch(line, -1)
			for _, funct := range match {
				switch funct[1] {
				case "window": // The window domain was called
					switch funct[2] { // Checks the function that were called from the domain
					case "x", "y":
						value, err := strconv.Atoi(funct[4])
						if err != nil {
							log.Fatalf("%s", err)
						}
						if funct[2] == "x" {
							window_setX(value)
						} else {
							window_setY(value)
						}
					case "title":
						window_setTitle(funct[4])
					case "url":
						window_setDst(funct[4])
					case "run":
						window_run()
					}

				case "system":
					fmt.Println(funct[2])

				default:
					notify_error("Unknown domain "+funct[2], "parser.interpeter()")
				}

			}
		}
	}
}

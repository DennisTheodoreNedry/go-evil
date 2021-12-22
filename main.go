package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	argument_add("--help", "-h", false, "Shows all available arguments and their purpose", []string{"NULL"})
	argument_add("--target_platform", "-tp", true, "For which platform should the malware be compiled for, options are [darwin, linux, windows]", []string{"darwin, linux, windws"})
	argument_add("--target_architecture", "-ta", true, "For which architecture should the malware be compiled for, options are [amd64, i386]", []string{"amd64", "i386"})
	argument_add("--file", "-f", true, "File to compile [REQUIRED]", []string{"NULL"})
	argument_add("--verbose", "-v", true, "How verbose should the program be, options are [1,2,3]", []string{"0", "1", "2", "3"})

	argument_parse() // Lets check what the user entered

	if len(os.Args[0:]) > 1 { // The user entered something
		if argument_check("-h") {
			argument_help()
		} else {
			var file = ""                      // Which file to compile
			var target_platform = runtime.GOOS // Default is the current system we are running on
			var architecture = runtime.GOARCH  // Default is the architecture we are currently running on

			if argument_check("-tp") { // The user specificed a target platform
				fmt.Println(argument_get("-tp"))
			}

			if !argument_check("-f") {
				notify_error("The '--file'/'-f' flag was not passed.", "main.main()")
			} else {
				file = argument_get("-f") // Get the file
			}

			if argument_check("-v") {
				verbose_lvl = argument_get("-v")
				notify_log("Setting verbose level to "+verbose_lvl, verbose_lvl, "1")
			}
			notify_log("File to compile is "+file, verbose_lvl, "1")
			notify_log("Malware will be compiled against "+target_platform, verbose_lvl, "2")
			notify_log("Malware will be compiled against a "+architecture+" architecture", verbose_lvl, "2")

			// Run interpreter
			interpeter(file)

			// Run compiler on the interpreted material
			write_file()   // The interpreter has filled the internal array with the correct go code, so this will dump it to a file
			compile_file() // This compiles the previously written code into a functioan program

		}
	} else {
		notify_error("No argument was provided, run '--help'/'-h' to have a look at the arguments available", "main.main()")
	}
}

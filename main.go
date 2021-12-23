package main

import (
	"fmt"
	"os"
	"runtime"

	arg "github.com/s9rA16Bf4/Malware_Language/utility/arguments"
	"github.com/s9rA16Bf4/Malware_Language/utility/io"
	"github.com/s9rA16Bf4/Malware_Language/utility/notify"
	"github.com/s9rA16Bf4/Malware_Language/utility/parser"
)

func main() {

	arg.Argument_add("--help", "-h", false, "Shows all available arguments and their purpose", []string{"NULL"})
	arg.Argument_add("--target_platform", "-tp", true, "For which platform should the malware be compiled for, options are [darwin, linux, windows]", []string{"darwin, linux, windws"})
	arg.Argument_add("--target_architecture", "-ta", true, "For which architecture should the malware be compiled for, options are [amd64, i386]", []string{"amd64", "i386"})
	arg.Argument_add("--file", "-f", true, "File to compile [REQUIRED]", []string{"NULL"})
	arg.Argument_add("--verbose", "-v", true, "How verbose should the program be, options are [1,2,3]", []string{"0", "1", "2", "3"})

	arg.Argument_parse() // Lets check what the user entered

	if len(os.Args[0:]) > 1 { // The user entered something
		if arg.Argument_check("-h") {
			arg.Argument_help()
		} else {
			var file = ""                      // Which file to compile
			var target_platform = runtime.GOOS // Default is the current system we are running on
			var architecture = runtime.GOARCH  // Default is the architecture we are currently running on

			if arg.Argument_check("-tp") { // The user specificed a target platform
				fmt.Println(arg.Argument_get("-tp"))
			}

			if !arg.Argument_check("-f") {
				notify.Notify_error("The '--file'/'-f' flag was not passed.", "main.main()")
			} else {
				file = arg.Argument_get("-f") // Get the file
			}

			if arg.Argument_check("-v") {
				notify.Verbose_lvl = arg.Argument_get("-v")
				notify.Notify_log("Setting verbose level to "+notify.Verbose_lvl, notify.Verbose_lvl, "1")
			}
			notify.Notify_log("File to compile is "+file, notify.Verbose_lvl, "1")
			notify.Notify_log("Malware will be compiled against "+target_platform, notify.Verbose_lvl, "2")
			notify.Notify_log("Malware will be compiled against a "+architecture+" architecture", notify.Verbose_lvl, "2")

			// Run interpreter
			parser.Interpeter(file)

			// Run compiler on the interpreted material
			io.Write_file()   // The interpreter has filled the internal array with the correct go code, so this will dump it to a file
			io.Compile_file() // This compiles the previously written code into a functioan program

		}
	} else {
		notify.Notify_error("No argument was provided, run '--help'/'-h' to have a look at the arguments available", "main.main()")
	}
}

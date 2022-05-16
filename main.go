package main

import (
	"fmt"
	"os"
	"runtime"

	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/json"
	"github.com/s9rA16Bf4/go-evil/utility/parser"
	"github.com/s9rA16Bf4/go-evil/utility/version"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func main() {
	arg.Argument_add("--help", "-h", false, "Shows all available arguments and their purpose", []string{"NULL"})
	arg.Argument_add("--target_platform", "-tp", true, "For which platform should the malware be compiled for", []string{"darwin", "linux", "windows"})
	arg.Argument_add("--target_architecture", "-ta", true, "For which architecture should the malware be compiled for", []string{"amd64", "i386"})
	arg.Argument_add("--file", "-f", true, "File to compile [REQUIRED]", []string{"NULL"})
	arg.Argument_add("--verbose", "-vv", true, "How verbose should the program be", []string{"0", "1", "2", "3"})
	arg.Argument_add("--debug", "-d", true, "Debug iptions", []string{"false", "true"})
	arg.Argument_add("--version", "-v", false, "Prints the compiler version", []string{"NULL"})
	arg.Argument_add("--output", "-o", true, "Name of the binary malware", []string{"NULL"})
	arg.Argument_add("--extension", "-exe", true, "Extension of the binary malware", []string{"NULL"})
	arg.Argument_add("--test_mode", "-tm", true, "Enables test mode on your malware, [THIS SHOULD NOT BE USED IN PRODUCTION]", []string{"false", "true"})
	arg.Argument_add("--interpreter", "-i", true, "A builtin interpreter that allows you to directly run your code", []string{"NULL"})
	arg.Argument_add("--exit_on_error", "-eoe", true, "Disables the malware from exiting if an error occurs. Default is false", []string{"true", "false"})
	arg.Argument_add("--print_json_data", "-pjd", false, "Prints the finalized json structure after compiling a file", []string{"NULL"})

	arg.Argument_parse() // Lets check what the user entered
	json_object := json.Create_object()
	json_object.Host_OS = runtime.GOOS
	json_object.Append_to_call("main")

	if len(os.Args[0:]) > 1 { // The user entered something
		if arg.Argument_check("-h") {
			arg.Argument_help()
		} else if arg.Argument_check("-v") {
			version.Print_version()
		} else if arg.Argument_check("-i") {
			parser.Interpreter(arg.Argument_get("-i"), json.Send(json_object))
		} else {
			if arg.Argument_check("-tp") { // The user specificed a target platform
				json_object.Target_OS = arg.Argument_get("-tp")
			} else {
				json_object.Target_OS = runtime.GOOS
			}
			if arg.Argument_check("-ta") {
				json_object.Target_ARCH = arg.Argument_get("-ta")
			} else {
				json_object.Target_ARCH = runtime.GOARCH
			}
			if !arg.Argument_check("-f") {
				notify.Error("The '--file'/'-f' flag was not passed.", "main.main()")
			} else {
				json_object.File = arg.Argument_get("-f") // Get and remember the file
			}
			if arg.Argument_check("-vv") {
				json_object.Verbose_LVL = arg.Argument_get("-vv")
				notify.Log("Setting verbose level to "+json_object.Verbose_LVL, json_object.Verbose_LVL, "1")
			} else {
				json_object.Verbose_LVL = "0"
			}

			if arg.Argument_check("-d") && arg.Argument_get("-d") == "true" {
				json_object.DebugMode = true
			} else {
				json_object.DebugMode = false
			}

			if arg.Argument_check("-o") {
				json_object.Binary_name = arg.Argument_get("-o")
			}
			if arg.Argument_check("-exe") {
				json_object.Binary_name = arg.Argument_get("-exe")
			}
			if arg.Argument_check("-tm") && arg.Argument_get("-tm") == "true" {
				json_object.TestMode = true
				notify.Log("Malware will be compiled in test mode", json_object.Verbose_LVL, "2")
			} else {
				json_object.TestMode = false
				notify.Log("Malware will be compiled in production mode", json_object.Verbose_LVL, "2")
			}

			if arg.Argument_check("-eoe") && arg.Argument_get("-eoe") == "true" {
				notify.Exit_on_error = true
			}

			notify.Log(fmt.Sprintf("File to compile is '%s'", json_object.File), json_object.Verbose_LVL, "1")
			notify.Log(fmt.Sprintf("Malware will be compiled against '%s'", json_object.Target_OS), json_object.Verbose_LVL, "2")
			notify.Log(fmt.Sprintf("Malware will be compiled against a '%s' architecture", json_object.Target_ARCH), json_object.Verbose_LVL, "2")

			// Run the parser
			base_64_serialize_json := parser.Parser(json.Send(json_object))
			data_structure := json.Receive(base_64_serialize_json)

			// Run compiler on the interpreted material
			data_structure = json.Receive(io.Write_file(json.Send(data_structure)))   // The interpreter has filled the internal array with the correct go code, so this will dump it to a file
			data_structure = json.Receive(io.Compile_file(json.Send(data_structure))) // This compiles the previously written code into a functioan program

			if arg.Argument_check("-pjd") { // These names are gonna give me a stroke
				fmt.Println(string(json.Convert_to_json(data_structure)))
			}
		}

	} else {
		notify.Error("No argument was provided, run '--help'/'-h' to have a look at the arguments available", "main.main()")
	}
}

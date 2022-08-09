package main

import (
	"fmt"
	"runtime"

	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/json"
	"github.com/s9rA16Bf4/go-evil/utility/parser"
	"github.com/s9rA16Bf4/go-evil/utility/version"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func main() {
	arg.Argument_add_with_options("--platform", "-p", true, "For which platform should the malware be compiled for", []string{"darwin", "linux", "windows"})
	arg.Argument_add_with_options("--architecture", "-a", true, "For which architecture should the malware be compiled for", []string{"amd64", "i386"})
	arg.Argument_add("--file", "-f", true, "File to compile [REQUIRED]")
	arg.Argument_add_with_options("--verbose", "-vv", true, "How verbose should the program be", []string{"0", "1", "2", "3"})
	arg.Argument_add_with_options("--debug", "-d", true, "Debug options, will not delete .go file after compile", []string{"false", "true"})
	arg.Argument_add("--version", "-v", false, "Prints the compiler version")
	arg.Argument_add("--output", "-o", true, "Name of the binary malware")
	arg.Argument_add("--extension", "-e", true, "Extension of the binary malware")
	arg.Argument_add("--json", "-j", false, "Prints the finalized json structure after compiling a file")

	parsed_flags := arg.Argument_parse() // Lets check what the user entered
	json_object := json.Create_object()
	json_object.Host_OS = runtime.GOOS
	json_object.Append_to_call("main")

	if len(parsed_flags) > 0 { // The user entered something
		if _, entered := parsed_flags["-v"]; entered {
			version.Print_version()

		} else {

			if value, entered := parsed_flags["-p"]; entered { // The user specificed a target platform
				json_object.Target_OS = value
			} else {
				json_object.Target_OS = runtime.GOOS
			}

			if value, entered := parsed_flags["-a"]; entered {
				json_object.Target_ARCH = value
			} else {
				json_object.Target_ARCH = runtime.GOARCH
			}

			if value, entered := parsed_flags["-f"]; entered {
				json_object.File = value // Get the file

			} else {
				notify.Error("The '--file'/'-f' flag was not passed.", "main.main()")
			}

			if value, entered := parsed_flags["-vv"]; entered {
				json_object.Verbose_LVL = value
				notify.Log(fmt.Sprintf("Setting verbose level to %s", json_object.Verbose_LVL), json_object.Verbose_LVL, "1")
			} else {
				json_object.Verbose_LVL = "0"
			}

			if value, entered := parsed_flags["-d"]; entered && (value == "True" || value == "true") {
				json_object.DebugMode = true
			} else {
				json_object.DebugMode = false
			}

			if value, entered := parsed_flags["-o"]; entered {
				json_object.Set_binary_name(value)
			}

			if value, entered := parsed_flags["-e"]; entered {
				json_object.Set_Extension(value)
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

			if _, entered := parsed_flags["-j"]; entered {
				fmt.Println(string(json.Convert_to_json(data_structure)))
			}
		}

	} else {
		notify.Error("No argument was provided, run '--help'/'-h' to have a look at the arguments available", "main.main()")
	}
}

package main

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/cleanup"
	"github.com/TeamPhoneix/go-evil/utility/json"
	"github.com/TeamPhoneix/go-evil/utility/parsing"
	"github.com/TeamPhoneix/go-evil/utility/version"

	"github.com/TeamPhoneix/go-evil/utility/io"
	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func main() {
	arg.Argument_add("--file", "-f", true, "File to compile [REQUIRED]")
	arg.Argument_add("--version", "-v", false, "Prints the compiler version")
	arg.Argument_add_with_options("--platform", "-p", true, "For which platform should the malware be compiled for", []string{"darwin", "linux", "windows"})
	arg.Argument_add_with_options("--architecture", "-a", true, "For which architecture should the malware be compiled for", []string{"amd64", "i386"})
	arg.Argument_add_with_options("--verbose", "-vv", true, "How verbose should the program be", []string{"0", "1", "2", "3"})
	arg.Argument_add_with_options("--debug", "-d", true, "Debug options, will not delete the src file after compilation", []string{"false", "true"})
	arg.Argument_add("--output", "-o", true, "Name of the binary malware")
	arg.Argument_add("--extension", "-e", true, "Extension of the binary malware")
	arg.Argument_add("--json", "-j", false, "Prints the finalized json structure after compiling a file")

	parsed := arg.Argument_parse()
	object := json.Create_object()

	if len(parsed) == 0 {
		notify.Error("No argument was provided", "main.main()")

	} else if _, ok := parsed["-v"]; ok {
		version.Version()

	} else {
		if value, ok := parsed["-f"]; ok {
			object.Set_file_path(value)
			notify.Inform(fmt.Sprintf("Compiling file %s", value))
		}

		if value, ok := parsed["-p"]; ok {
			object.Set_target_os(value)
		}

		if value, ok := parsed["-a"]; ok {
			object.Set_target_arch(value)
		}

		if value, ok := parsed["-vv"]; ok {
			object.Set_verbose_lvl(value)
		}

		if value, ok := parsed["-d"]; ok {
			object.Set_debug_mode(value)
		}

		if value, ok := parsed["-o"]; ok {
			object.Set_binary_name(value)
		}

		if value, ok := parsed["-e"]; ok {
			object.Set_extension(value)
		}

		if _, ok := parsed["-j"]; ok {
			object.Set_dump_json()
		}

		// Parse the file
		object = json.Receive(parsing.Parse(io.Read_file(json.Send(object)))) // Reads the file content and sends the result to the parser

		// Compile file
		object = json.Receive(io.Compile_file(json.Send(object)))

		// Cleanup
		cleanup.Start(json.Send(object))

	}
}

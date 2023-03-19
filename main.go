package main

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/version"
	"github.com/TeamPhoneix/go-evil/utility/wrapper"

	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func main() {
	arg.Argument_add("--file", "-f", true, "File to compile [REQUIRED]")
	arg.Argument_add("--version", "-v", false, "Prints the compiler version")
	arg.Argument_add_with_options("--platform", "-p", true, "For which platform should the malware be compiled for", []string{"darwin", "linux", "windows", "aix", "freebsd", "illumos", "js", "nacl", "netbsd", "openbasd", "plan9", "solaris"})
	arg.Argument_add_with_options("--architecture", "-a", true, "For which architecture should the malware be compiled for", []string{"amd64", "amd64p32", "386", "arm", "arm64", "ppc64", "pppc64le", "wasm", "mips", "mips64", "mips64le", "mipsle", "s390x"})
	arg.Argument_add_with_options("--verbose", "-vv", true, "How verbose should the program be", []string{"0", "1", "2", "3"})
	arg.Argument_add_with_options("--debug", "-d", true, "Debug options, will not delete the src file after compilation", []string{"false", "true"})
	arg.Argument_add("--output", "-o", true, "Name of the binary malware")
	arg.Argument_add("--extension", "-e", true, "Extension of the binary malware")
	arg.Argument_add("--json", "-j", false, "Prints the finalized json structure after compiling a file")
	arg.Argument_add("--obfuscate", "-ob", false, "Obfuscates the source code at compile time")
	arg.Argument_add_with_options("--debugger_behavior", "-db", true, "Changes the behavior of the malware after detecting a debugger", []string{"none", "stop", "remove", "loop"})
	arg.Argument_add("--build_directory", "-bd", true, "Sets the directory where all compiled files, source code, etc will be placed")

	parsed := arg.Argument_parse()
	object := structure.Create_json_object()

	if len(parsed) == 0 {
		notify.Error("No argument was provided", "main.main()")

	} else if _, ok := parsed["-v"]; ok {
		version.Version()

	} else {
		object.Set_verbose_lvl("0") // Default value

		for key, value := range parsed {
			switch key {
			case "-f":
				object.Set_file_path(value)
			case "-p":
				object.Set_target_os(value)
			case "-a":
				object.Set_target_arch(value)
			case "-vv":
				object.Set_verbose_lvl(value)
			case "-d":
				object.Set_debug_mode(value)
			case "-o":
				object.Set_binary_name(value)
			case "-e":
				object.Set_extension(value)
			case "-j":
				object.Set_dump_json()
			case "-ob":
				object.Enable_obfuscate()
			case "-db":
				object.Change_detection_behavior(value)

			case "-bd":
				object.Set_build_directory(value)
			}
		}

		notify.Log(fmt.Sprintf("Compiling file %s", object.File_path), object.Verbose_lvl, "1")

		wrapper.Parse_and_compile(structure.Send(object))
	}
}

package main

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/version"
	"github.com/s9rA16Bf4/go-evil/utility/wrapper"

	argumentparser "github.com/s9rA16Bf4/ArgumentParser"
)

func main() {
	object := structure.Create_json_object()

	handler := argumentparser.Constructor(true)
	handler.AddFunction("file", "f", true, true, "File to compile", object.Set_file_path)
	handler.AddFunction("version", "v", false, false, "Prints the compiler version", version.Version)
	handler.AddFunctionOptions("platform", "p", true, false, "For which platform should the malware be compiled for", object.Set_target_os, []string{"darwin", "linux", "windows", "aix", "freebsd", "illumos", "js", "nacl", "netbsd", "openbasd", "plan9", "solaris"})
	handler.AddFunctionOptions("architecture", "a", true, false, "For which architecture should the malware be compiled for", object.Set_target_arch, []string{"amd64", "amd64p32", "386", "arm", "arm64", "ppc64", "pppc64le", "wasm", "mips", "mips64", "mips64le", "mipsle", "s390x"})
	handler.AddFunctionOptions("verbose", "vv", true, false, "How verbose should the program be", object.Set_verbose_lvl, []string{"0", "1", "2", "3"})
	handler.AddFunctionOptions("debug", "d", true, false, "Debug options, will not delete the src file after compilation", object.Set_debug_mode, []string{"false", "true"})
	handler.AddFunction("output", "o", true, false, "Name of the binary malware", object.Set_binary_name)
	handler.AddFunction("extension", "e", true, false, "Extension of the binary malware", object.Set_extension)
	handler.AddFunction("dump_json", "dj", false, false, "Prints the finalized json structure after compiling a file", object.Set_dump_json)
	handler.AddFunction("obfuscate", "ob", false, false, "Obfuscates the source code at compile time", object.Enable_obfuscate)
	handler.AddFunctionOptions("debugger_behavior", "db", true, false, "Changes the behavior of the malware after detecting a debugger", object.Change_detection_behavior, []string{"none", "stop", "remove", "loop"})
	handler.AddFunction("build_directory", "bd", true, false, "Sets the directory where all compiled files, source code, etc will be placed", object.Set_build_directory)
	handler.AddFunction("alphabet", "A", true, false, "Sets the internal alphabet utilized", object.Set_alphabet)
	handler.AddFunction("external_domain_path", "xdp", true, false, "Adds an external domain which the program can use", object.Add_external_domains_path)

	handler.Parse()

	object.Log_object.Log(fmt.Sprintf("Compiling file %s", object.File_path), 1)

	wrapper.Parse_and_compile(&object)
}

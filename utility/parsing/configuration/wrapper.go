package configuration

import (
	"fmt"
	"regexp"

	evil_regex "github.com/s9rA16Bf4/go-evil/utility/parsing/regex"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
)

// Checks for a configuration section in the structure
func Check_configuration(data_object *json.Json_t) {

	regex := regexp.MustCompile(evil_regex.COMPILER_CONFIGURATION)
	result := regex.FindAllStringSubmatch(data_object.File_gut, -1)
	line := ""

	if len(result) == 0 {
		notify.Warning("Failed to find config section, will assume default values!")
	} else {
		line = result[0][1]
	}

	check_version(line) // Checks if this compiler can compile this version

	check_output(line, data_object)
	check_architecture(line, data_object)
	check_os(line, data_object)
	check_extension(line, data_object)
	check_obfuscate(line, data_object)
	check_debugger_behavior(line, data_object)

	data_object.Log_object.Log(fmt.Sprintf("The malware will be called `%s`", data_object.Binary_name), 2)
	data_object.Log_object.Log(fmt.Sprintf("The malware will be compiled for `%s`", data_object.Target_os), 2)
	data_object.Log_object.Log(fmt.Sprintf("The malware will be compiled for a/an `%s` based processor", data_object.Target_arch), 2)
	data_object.Log_object.Log(fmt.Sprintf("The malware will have the extension `%s`", data_object.Extension), 2)
	data_object.Log_object.Log(fmt.Sprintf("The malware will have the behavior `%s` if it's being run under a debugger", data_object.Debugger_behavior), 2)

	if data_object.Obfuscate {
		data_object.Log_object.Log("The binary file will be obfuscated", 2)
	} else {
		data_object.Log_object.Log("The binary file will not be obfuscated", 2)
	}

}

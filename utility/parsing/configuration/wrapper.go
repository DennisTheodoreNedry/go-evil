package configuration

import (
	"fmt"
	"regexp"

	evil_regex "github.com/TeamPhoneix/go-evil/utility/parsing/regex"
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Checks for a configuration section in the structure
func Check_configuration(s_json string) string {
	data_object := structure.Receive(s_json)

	regex := regexp.MustCompile(evil_regex.COMPILER_CONFIGURATION)
	result := regex.FindAllStringSubmatch(data_object.File_gut, -1)
	line := ""

	if len(result) == 0 {
		notify.Warning("Failed to find config section, will assume default values!")
	} else {
		line = result[0][1]
	}

	data_object = structure.Receive(check_output(line, structure.Send(data_object)))
	data_object = structure.Receive(check_architecture(line, structure.Send(data_object)))
	data_object = structure.Receive(check_os(line, structure.Send(data_object)))
	data_object = structure.Receive(check_extension(line, structure.Send(data_object)))
	data_object = structure.Receive(check_obfuscate(line, structure.Send(data_object)))
	data_object = structure.Receive(check_debugger_behavior(line, structure.Send(data_object)))

	notify.Log(fmt.Sprintf("The malware will be called `%s`", data_object.Binary_name), data_object.Verbose_lvl, "2")
	notify.Log(fmt.Sprintf("The malware will be compiled for `%s`", data_object.Target_os), data_object.Verbose_lvl, "2")
	notify.Log(fmt.Sprintf("The malware will be compiled for a/an `%s` based processor", data_object.Target_arch), data_object.Verbose_lvl, "2")
	notify.Log(fmt.Sprintf("The malware will have the extension `%s`", data_object.Extension), data_object.Verbose_lvl, "2")
	notify.Log(fmt.Sprintf("The malware will have the behavior `%s` if it's being run under a debugger", data_object.Debugger_behavior), data_object.Verbose_lvl, "2")

	if data_object.Obfuscate {
		notify.Log("The binary file will be obfuscated", data_object.Verbose_lvl, "2")
	} else {
		notify.Log("The binary file will not be obfuscated", data_object.Verbose_lvl, "2")
	}

	return structure.Send(data_object)
}

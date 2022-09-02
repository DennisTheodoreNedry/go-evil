package parsing

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Checks for a configuration section in the structure
//
//
func Check_configuration(s_json string) string {
	data_object := structure.Receive(s_json)

	regex := regexp.MustCompile(COMPILER_CONFIGURATION)
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

	notify.Log(fmt.Sprintf("The malware will be called `%s`", data_object.Binary_name), data_object.Verbose_lvl, "2")
	notify.Log(fmt.Sprintf("The malware will be compiled for `%s`", data_object.Target_os), data_object.Verbose_lvl, "2")
	notify.Log(fmt.Sprintf("The malware will be for a/n `%s` based processor", data_object.Target_arch), data_object.Verbose_lvl, "2")
	notify.Log(fmt.Sprintf("The malware will have the extension `%s`", data_object.Extension), data_object.Verbose_lvl, "2")

	if data_object.Obfuscate {
		notify.Log("The source code will be obfuscated", data_object.Verbose_lvl, "2")
	} else {
		notify.Log("The source code will not be obfuscated", data_object.Verbose_lvl, "2")
	}

	return structure.Send(data_object)
}

//
//
// Checks if the user specificed a name for the malware
//
//
func check_output(line string, s_json string) string {
	data_object := structure.Receive(s_json)

	if data_object.Binary_name == "" { // Don't override if the user already have provided a name
		regex := regexp.MustCompile(CONFIGURATION_NAME)
		result := regex.FindAllStringSubmatch(line, -1)
		name := "me_not_a_virus"

		if len(result) > 0 {
			name = result[0][1]
		}

		data_object.Set_binary_name(name)
	}

	return structure.Send(data_object)
}

//
//
// Checks if the user specificed an architecture for the malware
//
//
func check_architecture(line string, s_json string) string {
	data_object := structure.Receive(s_json)

	if data_object.Target_arch == "" { // Don't override if the user already have provided a name
		regex := regexp.MustCompile(CONFIGURATION_ARCH)
		result := regex.FindAllStringSubmatch(line, -1)
		arch := runtime.GOARCH

		if len(result) > 0 {
			arch = result[0][1]
		}

		data_object.Set_target_arch(arch)

	}

	return structure.Send(data_object)
}

//
//
// Checks if the user specificed an os for the malware
//
//
func check_os(line string, s_json string) string {
	data_object := structure.Receive(s_json)

	if data_object.Target_os == "" { // Don't override if the user already have provided a os
		regex := regexp.MustCompile(CONFIGURATION_OS)
		result := regex.FindAllStringSubmatch(line, -1)
		os := runtime.GOOS

		if len(result) > 0 {
			os = result[0][1]
		}

		data_object.Set_target_os(os)

	}

	return structure.Send(data_object)
}

//
//
// Checks if the user specificed an extension for the malware
//
//
func check_extension(line string, s_json string) string {
	data_object := structure.Receive(s_json)

	if data_object.Extension == "" { // Don't override if the user already have provided a extension
		regex := regexp.MustCompile(CONFIGURATION_EXTENSION)
		result := regex.FindAllStringSubmatch(line, -1)
		ext := ""

		if len(result) > 0 {
			ext = result[0][1]

			check := tools.Starts_with(ext, []string{"."}) // Checks if the extension starts with a dot

			if status := check["."]; !status { // It doesn't
				ext = fmt.Sprintf(".%s", ext)
			}

		} else if os.Getenv("GOOS") == "windows" {
			ext = ".exe"
		}

		data_object.Set_extension(ext)
	}

	return structure.Send(data_object)
}

//
//
// Checks if the user specificed the output to be obfuscated or not
//
//
func check_obfuscate(line string, s_json string) string {
	data_object := structure.Receive(s_json)

	if !data_object.Obfuscate { // No point in checking if the user already has enabled it
		regex := regexp.MustCompile(CONFIGURATION_OBFUSCATE)
		result := regex.FindAllStringSubmatch(line, -1)

		if len(result) > 0 {
			option := strings.ToLower(result[0][1])
			if option == "true" {
				data_object.Enable_obfuscate()
			}
		}
	}

	return structure.Send(data_object)
}

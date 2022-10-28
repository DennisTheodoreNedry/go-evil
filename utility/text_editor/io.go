package text_editor

import (
	"io/ioutil"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/wrapper"
)

//
//
// Wrapper modifying the json structure and sending it of to the core wrapper
//
//
func Compile(s_json string, name string, os string, arch string, obf string, ext string) {
	data_object := structure.Receive(s_json)

	if name != "" {
		data_object.Set_binary_name(name)
	}

	if os != "default" {
		data_object.Set_target_os(os)
	}

	if arch != "default" {
		data_object.Set_target_arch(arch)
	}

	if obf == "yes" {
		data_object.Enable_obfuscate()
	} else {
		data_object.Disable_obfuscate()
	}

	if ext != "" {
		data_object.Set_extension(ext)
	}

	wrapper.Parse_and_compile(structure.Send(data_object))
}

//
//
// Writes the contents entered in the text editor to the disk
//
//
func Save(content string, path string) {
	ioutil.WriteFile(path, []byte(content), 0644)
}

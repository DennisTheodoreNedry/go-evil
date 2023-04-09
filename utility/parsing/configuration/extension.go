package configuration

import (
	"os"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Checks if the user specificed an extension for the malware
func check_extension(line string, s_json string) string {
	data_object := structure.Receive(s_json)

	if data_object.Extension == "" { // Don't override if the user already have provided a extension
		ext := ""

		if os.Getenv("GOOS") == "windows" {
			ext = ".exe"
		}

		data_object.Set_extension(ext)
	}

	return structure.Send(data_object)
}

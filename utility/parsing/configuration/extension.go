package configuration

import (
	"os"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Checks if the user specificed an extension for the malware
func check_extension(line string, data_object *json.Json_t) {

	if data_object.Extension == "" { // Don't override if the user already have provided a extension
		ext := ""

		if os.Getenv("GOOS") == "windows" {
			ext = ".exe"
		}

		data_object.Set_extension(ext)
	}

}

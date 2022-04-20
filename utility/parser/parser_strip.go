package parser

import (
	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Strip(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("Parser.Strip()")

	// Removes everything until the line after main:{
	eof := false
	main := false
	var new_file_gut []string

	for _, funct := range data_structure.File_gut {

		if contains.StartsWith(funct, []string{"}"}) {
			eof = true
		} else if contains.StartsWith(funct, []string{"main"}) {
			main = true
			continue
		}

		if main && !eof {
			for i, c := range funct {
				if string(c) == "@" { // Removes comments

					notify.Log("Found comment, will ignore line", data_structure.Verbose_LVL, "3")

					funct = funct[0:i]
					break
				}
			}
			new_file_gut = append(new_file_gut, funct)
		}

		if eof {
			break
		}
	}

	data_structure.File_gut = new_file_gut // Updated array
	data_structure.File_count = len(new_file_gut)

	return json.Send(data_structure)
}

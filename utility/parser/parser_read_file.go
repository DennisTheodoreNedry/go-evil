package parser

import (
	"strings"

	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

func Read_file(base_64_serialize_json string) string {

	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("Parser.Read_file()")

	content := io.Read_file(data_structure.File)
	split_content := strings.Split(content, "\n")

	for _, line := range split_content {
		trimmed := strings.Trim(line, "\t ")
		if trimmed != "" && string(trimmed[0]) == "@" {
			continue
		}
		data_structure.Append_file_gut(trimmed)
	}

	return json.Send(data_structure)
}

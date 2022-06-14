package parser

import (
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_IMPORTS = "use +([a-zA-Z]+);" // Extracts the import
)

func Imports(base_64_serialize_json string) string {

	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("Parser.Imports()")

	regex := regexp.MustCompile(EXTRACT_IMPORTS)

	for _, line := range data_structure.File_gut {
		result := regex.FindAllStringSubmatch(line, -1)
		if len(result) > 0 {
			data_structure.Append_imported_domain(result[0][1])
		}
	}

	return json.Send(data_structure)
}

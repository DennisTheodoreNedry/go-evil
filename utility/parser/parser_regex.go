package parser

import (
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_DOMAIN = "([a-zA-Z#]+)\\." // Extracts the domain
	//EXTRACT_IMPORTS = ""
)

func Regex(base_64_serialize_json string) string {

	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("Parser.Regex()")

	regex := regexp.MustCompile(EXTRACT_DOMAIN)

	for _, line := range data_structure.File_gut {
		result := regex.FindAllStringSubmatch(line, -1)
		data_structure.Append_File_domain(result[0][1])

	}

	return json.Send(data_structure)
}

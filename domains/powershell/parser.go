package powershell

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(powershell|#pwsh)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(powershell|#pwsh)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("powershell.Parse()")

	for _, line := range data_structure.File_gut {

		value := contains.Passed_value(line)

		regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
		result := regex.FindAllStringSubmatch(line, -1)

		if len(result) > 0 { // There is a subdomain to extract
			subdomain := result[0][2]
			function := result[0][3]

			switch subdomain {
			case "disable":
				switch function {
				case "defender":
					data_structure.Append_malware_gut("pwsh.Disable_defender()")
				default:
					error.Function_error(function, "powershell.Parse()")
				}
			case "change":
				switch function {
				case "wallpaper":
					data_structure.Append_malware_gut(fmt.Sprintf("pwsh.Change_wallpaper(\"%s\")", value))
				default:
					error.Function_error(function, "powershell.Parse()")
				}
			case "set":
				switch function {
				case "execution_policy":
					data_structure.Append_malware_gut(fmt.Sprintf("pwsh.Execution_Policy(\"%s\")", value))
				default:
					error.Function_error(function, "powershell.Parse()")
				}

			default:
				error.Subdomain_error(subdomain, "powershell.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][1]
				switch function {
				default:
					error.Function_error(function, "powershell.Parse()")
				}
			}
		}
	}

	return json.Send(data_structure)
}

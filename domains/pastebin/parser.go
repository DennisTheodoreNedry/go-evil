package pastebin

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(pastebin|#paste)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(pastebin|#paste)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("pastebin.Parse()")

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
				case "username":
					data_structure.Append_malware_gut(fmt.Sprintf("pastebin.Set_username(\"%s\")", value))
				case "password":
					data_structure.Append_malware_gut(fmt.Sprintf("pastebin.Set_password(\"%s\")", value))
				case "token":
					data_structure.Append_malware_gut(fmt.Sprintf("pastebin.Set_token(\"%s\")", value))
				case "content":
					data_structure.Append_malware_gut(fmt.Sprintf("pastebin.Set_content(\"%s\")", value))
				case "titel":
					data_structure.Append_malware_gut(fmt.Sprintf("pastebin.Set_titel(\"%s\")", value))
				case "expiration_time":
					data_structure.Append_malware_gut(fmt.Sprintf("pastebin.Set_expiration(\"%s\")", value))
				case "visibility":
					data_structure.Append_malware_gut(fmt.Sprintf("pastebin.Set_visibility(\"%s\")", value))
				case "key":
					data_structure.Append_malware_gut(fmt.Sprintf("pastebin.Set_key(\"%s\")", value))

				default:
					error.Function_error(function, "pastebin.Parse()")
				}
			default:
				error.Subdomain_error(subdomain, "pastebin.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]
				switch function {
				case "paste":
					data_structure.Append_malware_gut("pastebin.Paste()")
				default:
					error.Function_error(function, "pastebin.Parse()")
				}
			}
		}
	}

	return json.Send(data_structure)
}

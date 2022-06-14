package time

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(time|#wait)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(time|#wait)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("time.Parse()")

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
				case "year":
					data_structure.Append_malware_gut(fmt.Sprintf("time.SetYear(\"%s\")", value))
				case "month":
					data_structure.Append_malware_gut(fmt.Sprintf("time.SetMonth(\"%s\")", value))
				case "day":
					data_structure.Append_malware_gut(fmt.Sprintf("time.SetDay(\"%s\")", value))
				case "hour":
					data_structure.Append_malware_gut(fmt.Sprintf("time.SetHour(\"%s\")", value))
				case "min":
					data_structure.Append_malware_gut(fmt.Sprintf("time.SetMin(\"%s\")", value))
				default:
					error.Function_error(function, "time.Parse()")
				}
			default:
				error.Subdomain_error(subdomain, "time.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]
				switch function {
				case "run":
					data_structure.Append_malware_gut("time.Run()")
				case "until":
					data_structure.Append_malware_gut(fmt.Sprintf("time.Until(\"%s\")", value))
				default:
					error.Function_error(function, "time.Parse()")
				}
			}
		}
	}
	return json.Send(data_structure)

}

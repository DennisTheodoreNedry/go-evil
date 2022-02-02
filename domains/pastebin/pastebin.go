package pastebin

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN      = "(pastebin|#paste)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION_VALUE = ".+\\(\"(.*)\"\\);"                        // Grabs the value being passed to the function
	EXTRACT_FUNCTION       = "(pastebin|#paste)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
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
			case "username":
				mal.AddContent("pastebin.Set_username(\"" + value + "\")")
			case "password":
				mal.AddContent("pastebin.Set_password(\"" + value + "\")")
			case "token":
				mal.AddContent("pastebin.Set_token(\"" + value + "\")")
			case "content":
				mal.AddContent("pastebin.Set_content(\"" + value + "\")")
			case "titel":
				mal.AddContent("pastebin.Set_titel(\"" + value + "\")")
			case "expiration_time":
			case "visibility":
			case "key":
				mal.AddContent("pastebin.Set_key(\"" + value + "\")")

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
			case "paste":
				mal.AddContent("pastebin.Paste()")
			default:
				function_error(function)
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain '"+subdomain+"'", "pastebin.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function '"+function+"'", "pastebin.Parse()")
}

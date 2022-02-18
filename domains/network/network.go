package network

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN      = "(network|#net)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION_VALUE = ".+\\(\"(.*)\"\\);"                     // Grabs the value being passed to the function
	EXTRACT_FUNCTION       = "(network|#net)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
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
		case "get":
			switch function {
			case "save_disk":
				mal.AddContent("net.GET_save_disk()")
			case "save_variable":
				mal.AddContent("net.GET_save_variable()")
			case "set_prefix":
				mal.AddContent("net.GET_set_prefix(\"" + value + "\")")
			default:
				function_error(function)
			}
		case "post":
			switch function {
			case "add_header":
				mal.AddContent("net.POST_add_header(\"" + value + "\")")
			case "set_header":
				mal.AddContent("net.POST_set_header(\"" + value + "\")")
			case "bind_value":
				mal.AddContent("net.POST_bind_value_to_latest_header(\"" + value + "\")")
			default:
				function_error(function)
			}
		case "ping":
			switch function {
			case "set_max":
				mal.AddContent("net.Ping_set_roof(\"" + value + "\")")
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
			case "post":
				mal.AddContent("net.POST(\"" + value + "\")")
			case "get":
				mal.AddContent("net.GET(\"" + value + "\")")
			case "ping":
				mal.AddContent("net.Ping(\"" + value + "\")")
			default:
				function_error(function)
			}
		}
	}
}

func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain '"+subdomain+"'", "network.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function '"+function+"'", "network.Parse()")
}

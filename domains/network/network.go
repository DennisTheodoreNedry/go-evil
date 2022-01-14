package network

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN = "[a-z]+\\.([a-z]+)\\.([a-z]+)\\(\"(.*)\"\\);"
	EXTRACT_FUNCTION  = "(network|#net)\\.([a-z]+)\\((\"(.+)\")?\\);"
)

func Parse(new_line string) {
	regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
	result := regex.FindAllStringSubmatch(new_line, -1)
	if len(result) > 0 { // There is a subdomain to extract
		switch result[0][1] {
		case "get":
			switch result[0][2] {
			case "save_disk":
				mal.AddContent("net.GET_save_disk()")
			case "save_variable":
				mal.AddContent("net.GET_save_variable()")
			case "set_prefix":
				mal.AddContent("net.GET_set_prefix(\"" + result[0][4] + "\")")
			default:
				function_error(result[0][2])
			}
		case "post":
			switch result[0][2] {
			case "add_header":
				mal.AddContent("net.POST_add_header(\"" + result[0][4] + "\")")
			case "set_header":
				mal.AddContent("net.POST_set_header(\"" + result[0][4] + "\")")
			case "bind_value":
				mal.AddContent("net.POST_bind_value_to_latest_header(\"" + result[0][4] + "\")")
			default:
				function_error(result[0][2])
			}
		case "ping":
			switch result[0][2] {
			case "set_max":
				mal.AddContent("net.Ping_set_roof(\"" + result[0][4] + "\")")
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
			case "post":
				mal.AddContent("net.POST(\"" + result[0][3] + "\")")
			case "get":
				mal.AddContent("net.GET(\"" + result[0][3] + "\")")
			case "ping":
				mal.AddContent("net.Ping(\"" + result[0][3] + "\")")
			default:
				function_error(result[0][1])
			}
		}
	}
}

func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain "+subdomain, "network.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function "+function, "network.Parse()")
}

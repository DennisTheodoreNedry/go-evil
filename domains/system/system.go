package system

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN = "[a-z]+\\.([a-z]+)\\.([a-z]+)\\(\"(.*)\"\\);"
	EXTRACT_FUNCTION  = "system\\.([a-z]+)\\((\"(.+)\")?\\);"
)

func Parse(new_line string) {
	regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
	result := regex.FindAllStringSubmatch(new_line, -1)
	if len(result) > 0 { // There is a subdomain to extract
		switch result[0][1] {
		default:
			subdomain_error(result[0][1])
		}
	} else { // There might be a function which doesn't require a subdomain to work
		regex := regexp.MustCompile(EXTRACT_FUNCTION)
		result := regex.FindAllStringSubmatch(new_line, -1)
		if len(result) > 0 {
			switch result[0][1] {
			case "exit":
				mal.AddContent("sys.System_exit(\"" + result[0][3] + "\")")
			case "out":
				mal.AddContent("sys.System_out(\"" + result[0][3] + "\")")
			case "add_to_startup":
				mal.AddContent("sys.add_to_startup()")
			case "spawn":
				io.Append_domain("syscall") // Needed
				mal.AddContent("syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)")
			case "in":
				mal.AddContent("sys.User_input()")
			default:
				function_error(result[0][1])
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain "+subdomain, "system.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function "+function, "system.Parse()")
}

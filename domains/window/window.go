package window

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN = "[a-z]+\\.([a-z]+)\\.([a-z]+)\\(\"(.*)\"\\);"
	EXTRACT_FUNCTION  = "window\\.([a-z]+)\\((\"(.+)\")?\\);"
)

func Parse(new_line string) {
	regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
	result := regex.FindAllStringSubmatch(new_line, -1)
	if len(result) > 0 { // There is a subdomain to extract
		switch result[0][1] {
		case "set":
			switch result[0][2] {
			case "x":
				mal.AddContent("win.SetX(\"" + result[0][3] + "\")")
			case "y":
				mal.AddContent("win.SetY(\"" + result[0][3] + "\")")
			case "title":
				mal.AddContent("win.SetTitle(\"" + result[0][3] + "\")")
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
			case "goto":
				mal.AddContent("win.GoToUrl(\"" + result[0][3] + "\")")
			case "display":
				mal.AddContent("win.Display(\"" + result[0][3] + "\")")
			default:
				function_error(result[0][1])
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain "+subdomain, "window.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function "+function, "window.Parse()")
}

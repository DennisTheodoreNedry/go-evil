package time

import (
	"regexp"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN = "[a-z]+\\.([a-z]+)\\.([a-z]+)\\(\"(.*)\"\\);"
	EXTRACT_FUNCTION  = "(time|#wait)\\.([a-z]+)\\((\"(.+)\")?\\);"
)

func Parse(new_line string) {
	regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
	result := regex.FindAllStringSubmatch(new_line, -1)
	if len(result) > 0 { // There is a subdomain to extract
		switch result[0][1] {
		case "set":
			switch result[0][2] {
			case "year":
				mal.AddContent("time.SetYear(\"" + result[0][4] + "\")")
			case "month":
				mal.AddContent("time.SetMonth(\"" + result[0][4] + "\")")
			case "day":
				mal.AddContent("time.SetDay(\"" + result[0][4] + "\")")
			case "hour":
				mal.AddContent("time.SetHour(\"" + result[0][4] + "\")")
			case "min":
				mal.AddContent("time.SetMin(\"" + result[0][4] + "\")")
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
			case "run":
				mal.AddContent("time.Run()")
			case "until":
				mal.AddContent("time.Until(\"" + result[0][3] + "\")")
			default:
				function_error(result[0][1])
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain "+subdomain, "time.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function "+function, "time.Parse()")
}

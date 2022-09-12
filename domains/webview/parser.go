package webview

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// The main parser for the time domain
//
//
func Parser(function string, value string, s_json string) (string, string) {
	call := ""

	switch function {
	case "run":
		call, s_json = run(s_json)

	case "html":
		s_json = set_html(value, s_json)

	case "js":
		s_json = set_js(value, s_json)

	case "css":
		s_json = set_css(value, s_json)

	case "width":
		s_json = set_width(value, s_json)

	case "height":
		s_json = set_height(value, s_json)

	case "title":
		s_json = set_title(value, s_json)

	case "bind":
		s_json = bind(value, s_json)

	case "navigate":
		call, s_json = navigate(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call, s_json
}

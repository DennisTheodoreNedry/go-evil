package window

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/window/bind"
	"github.com/TeamPhoneix/go-evil/domains/window/navigate"
	"github.com/TeamPhoneix/go-evil/domains/window/run"
	"github.com/TeamPhoneix/go-evil/domains/window/set"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// The main parser for the window domain
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "run":
		call, s_json = run.Run(s_json)

	case "html":
		s_json = set.HTML(value, s_json)

	case "js":
		s_json = set.Javascript(value, s_json)

	case "css":
		s_json = set.Css(value, s_json)

	case "width":
		s_json = set.Width(value, s_json)

	case "height":
		s_json = set.Height(value, s_json)

	case "title":
		s_json = set.Title(value, s_json)

	case "bind":
		s_json = bind.Bind(value, s_json)

	case "navigate":
		call, s_json = navigate.Navigate(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "window.Parser()")

	}

	return call, s_json
}

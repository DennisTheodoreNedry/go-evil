package window

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/window/bind"
	"github.com/s9rA16Bf4/go-evil/domains/window/navigate"
	"github.com/s9rA16Bf4/go-evil/domains/window/run"
	"github.com/s9rA16Bf4/go-evil/domains/window/set"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "run":
		call = run.Run(data_object)

	case "html":
		set.HTML(value, data_object)

	case "js":
		set.Javascript(value, data_object)

	case "css":
		set.Css(value, data_object)

	case "width":
		set.Width(value, data_object)

	case "height":
		set.Height(value, data_object)

	case "title":
		set.Title(value, data_object)

	case "bind":
		bind.Bind(value, data_object)

	case "navigate":
		call = navigate.Navigate(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "window.Parser()")

	}

	return call
}

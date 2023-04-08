package json

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Adds a binding to the window
// it's accessible by running `window.<js_call>()` in your html code
func (object *Json_t) Add_binding(js_call string, evil_call string) {
	evil_call = tools.Erase_delimiter(evil_call, []string{"\""}, -1)

	object.Bind_gut[fmt.Sprintf("\"%s\"", js_call)] = evil_call
}

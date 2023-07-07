package json

import (
	"fmt"

	tools "github.com/s9rA16Bf4/Go-tools"
)

// Adds a binding to the window
// it's accessible by running `window.<js_call>()` in your html code
func (object *Json_t) Add_binding(js_call string, evil_call string) {
	evil_call = tools.EraseDelimiter(evil_call, []string{"\""}, -1)

	object.Bind_gut[fmt.Sprintf("\"%s\"", js_call)] = evil_call
}

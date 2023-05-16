package json

import "github.com/s9rA16Bf4/go-evil/utility/tools"

// Sets the js code being used in the window
func (object *Json_t) Set_js(content string) {
	content = tools.Erase_delimiter(content, []string{"\""}, -1)
	object.Js_gut = append(object.Js_gut, content)
}

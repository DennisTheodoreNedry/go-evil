package json

import "github.com/s9rA16Bf4/Go-tools/tools"

// Sets the css code being used in the window
func (object *Json_t) Set_css(content string) {
	content = tools.Erase_delimiter(content, []string{"\""}, -1)
	object.Css_gut = append(object.Css_gut, content)
}

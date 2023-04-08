package json

import "github.com/TeamPhoneix/go-evil/utility/tools"

// Sets the css code being used in the window
func (object *Json_t) Set_css(content string) {
	content = tools.Erase_delimiter(content, []string{"\""}, -1)
	object.Css_gut = append(object.Css_gut, content)
}

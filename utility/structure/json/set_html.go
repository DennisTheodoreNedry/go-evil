package json

import "github.com/TeamPhoneix/go-evil/utility/tools"

// Sets the html code being displayed
func (object *Json_t) Set_html(content string) {
	content = tools.Erase_delimiter(content, []string{"\""}, -1)
	object.Html_gut = append(object.Html_gut, content)
}

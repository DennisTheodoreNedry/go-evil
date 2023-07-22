package json

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
)

// Sets the html code being displayed
func (object *Json_t) Set_html(content string) {
	content = gotools.EraseDelimiter(content, []string{"\""}, -1)
	object.Html_gut = append(object.Html_gut, content)
}

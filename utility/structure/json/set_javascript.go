package json

import gotools "github.com/DennisTheodoreNedry/Go-tools"

// Sets the js code being used in the window
func (object *Json_t) Set_js(content string) {
	content = gotools.EraseDelimiter(content, []string{"\""}, -1)
	object.Js_gut = append(object.Js_gut, content)
}

package json

import gotools "github.com/s9rA16Bf4/Go-tools"

// Sets the width of the text editor
func (object *Json_t) Set_width(value string) {
	if i_value := gotools.StringToInt(value); i_value != -1 {
		object.Width = i_value
	} else {
		object.Width = 600
	}
}

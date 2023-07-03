package json

import "github.com/s9rA16Bf4/Go-tools/tools"

// Sets the width of the text editor
func (object *Json_t) Set_width(value string) {
	if i_value := tools.String_to_int(value); i_value != -1 {
		object.Width = i_value
	} else {
		object.Width = 600
	}
}

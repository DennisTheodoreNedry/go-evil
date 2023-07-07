package json

import tools "github.com/s9rA16Bf4/Go-tools"

// Sets the height of the text editor
func (object *Json_t) Set_height(value string) {
	if i_value := tools.StringToInt(value); i_value != -1 {
		object.Height = i_value
	} else {
		object.Height = 800
	}
}

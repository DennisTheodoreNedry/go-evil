package json

import "github.com/TeamPhoneix/go-evil/utility/tools"

// Sets the height of the text editor
func (object *Json_t) Set_height(value string) {
	if i_value := tools.String_to_int(value); i_value != -1 {
		object.Height = i_value
	} else {
		object.Height = 800
	}
}

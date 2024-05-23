package json

import gotools "github.com/DennisTheodoreNedry/Go-tools"

// Sets the height of the text editor
func (object *Json_t) Set_height(value string) {
	if i_value := gotools.StringToInt(value); i_value != -1 {
		object.Height = i_value
	} else {
		object.Height = 800
	}
}

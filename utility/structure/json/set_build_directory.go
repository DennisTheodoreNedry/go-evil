package json

import "github.com/s9rA16Bf4/Go-tools/tools"

// Set the build directory
func (object *Json_t) Set_build_directory(new_bd string) {
	if ok := tools.Ends_with(new_bd, []string{"/"})[new_bd]; !ok {
		new_bd += "/"
	}

	object.Build_directory = new_bd
}

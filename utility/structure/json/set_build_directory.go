package json

import tools "github.com/s9rA16Bf4/Go-tools"

// Set the build directory
func (object *Json_t) Set_build_directory(new_bd string) string {
	if ok := tools.EndsWith(new_bd, []string{"/"})[new_bd]; !ok {
		new_bd += "/"
	}

	object.Build_directory = new_bd

	return ""
}

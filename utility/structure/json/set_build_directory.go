package json

import gotools "github.com/DennisTheodoreNedry/Go-tools"

// Set the build directory
func (object *Json_t) Set_build_directory(new_bd string) string {
	if ok := gotools.EndsWith(new_bd, []string{"/"})[new_bd]; !ok {
		new_bd += "/"
	}

	object.Build_directory = new_bd

	return ""
}

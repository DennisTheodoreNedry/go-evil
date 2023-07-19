package json

import (
	"fmt"
	"path"
	"strings"

	gotools "github.com/s9rA16Bf4/Go-tools"
)

// Sets the binaries name
func (object *Json_t) Set_binary_name(name string) string {

	object.Malware_path = fmt.Sprintf("%s%s/", object.Build_directory, path.Dir(name))

	name = path.Base(name)
	result := gotools.Contains(name, []string{"."})

	if ok := result["."]; ok { // It contains a dot which we are gonna interpreted as an extension
		split := strings.Split(name, ".")
		name = split[0]
		object.Set_extension(split[1])
	}

	object.Binary_name = name
	return ""
}

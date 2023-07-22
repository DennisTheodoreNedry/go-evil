package json

import (
	"fmt"

	gotools "github.com/s9rA16Bf4/Go-tools"
)

// Sets the the extension
func (object *Json_t) Set_extension(ext string) string {

	if object.Extension == "" { // Update only if it doesn't contain anything
		result := gotools.Contains(ext, []string{"."}) // Checks if the extension contains a dot

		if status := result["."]; !status && ext != "" {
			ext = fmt.Sprintf(".%s", ext)
		}

		object.Extension = ext
	}
	return ""
}

package json

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Sets the the extension
func (object *Json_t) Set_extension(ext string) {

	if object.Extension == "" { // Update only if it doesn't contain anything
		result := tools.Contains(ext, []string{"."}) // Checks if the extension contains a dot

		if status := result["."]; !status && ext != "" {
			ext = fmt.Sprintf(".%s", ext)
		}

		object.Extension = ext
	}

}

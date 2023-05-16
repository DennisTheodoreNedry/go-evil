package set

import (
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
)

// Sets the height of the window
func Height(new_height string, data_object *json.Json_t) {
	new_height = tools.Erase_delimiter(new_height, []string{"\""}, -1)

	data_object.Set_height(new_height)

}

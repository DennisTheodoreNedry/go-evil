package set

import (
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
)

// Sets the width of the window
func Width(new_width string, data_object *json.Json_t) {
	new_width = tools.Erase_delimiter(new_width, []string{"\""}, -1)

	data_object.Set_width(new_width)

}

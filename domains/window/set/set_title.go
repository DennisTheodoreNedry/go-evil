package set

import (
	"github.com/TeamPhoneix/go-evil/utility/structure/json"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Sets the title of the window that appears
func Title(new_title string, data_object *json.Json_t) {
	new_title = tools.Erase_delimiter(new_title, []string{"\""}, -1)

	data_object.Set_title(new_title)

}

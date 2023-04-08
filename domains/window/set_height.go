package window

import (
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Sets the height of the window
func set_height(new_height string, s_json string) string {
	data_object := structure.Receive(s_json)
	new_height = tools.Erase_delimiter(new_height, []string{"\""}, -1)

	data_object.Set_height(new_height)

	return structure.Send(data_object)
}

package window

import (
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Sets the width of the window
func set_width(new_width string, s_json string) string {
	data_object := structure.Receive(s_json)
	new_width = tools.Erase_delimiter(new_width, []string{"\""}, -1)

	data_object.Set_width(new_width)

	return structure.Send(data_object)
}

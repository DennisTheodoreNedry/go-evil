package set

import (
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Sets the title of the window that appears
func Title(new_title string, s_json string) string {
	data_object := structure.Receive(s_json)
	new_title = tools.Erase_delimiter(new_title, []string{"\""}, -1)

	data_object.Set_title(new_title)

	return structure.Send(data_object)
}

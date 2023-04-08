package window

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Sets the css that will be used
func set_css(css_content string, s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Set_css(css_content)

	return structure.Send(data_object)
}

package window

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Sets the html content displayed
func set_html(html_content string, s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Set_html(html_content)

	return structure.Send(data_object)
}

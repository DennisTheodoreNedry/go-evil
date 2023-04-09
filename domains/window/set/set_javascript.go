package set

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Sets the js that wil be used
func Javascript(js_content string, s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Set_js(js_content)

	return structure.Send(data_object)
}

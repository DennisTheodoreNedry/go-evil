package set

import (
	"github.com/TeamPhoneix/go-evil/utility/structure/json"
)

// Sets the html content displayed
func HTML(html_content string, data_object *json.Json_t) {

	data_object.Set_html(html_content)

}

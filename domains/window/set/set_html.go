package set

import (
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Sets the html content displayed
func HTML(html_content string, data_object *json.Json_t) {

	data_object.Set_html(html_content)

}

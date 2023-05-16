package set

import (
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Sets the css that will be used
func Css(css_content string, data_object *json.Json_t) {

	data_object.Set_css(css_content)

}

package bind

import (
	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Binds a go function to a corresponding javascript function
func Bind(values string, data_object *json.Json_t) {
	arr := structure.Create_evil_object(values)

	js_call := arr.Get(0)
	evil_func := arr.Get(1)

	data_object.Add_binding(js_call, evil_func)

}

package strip

import "github.com/s9rA16Bf4/go-evil/utility/structure/json"

func Strip(data_object *json.Json_t) {
	remove_comments(data_object)
	remove_configuration(data_object)
	remove_imports(data_object)
	remove_injected_headers(data_object)
	remove_injected_code(data_object)

}

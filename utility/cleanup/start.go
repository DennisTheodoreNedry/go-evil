package cleanup

import "github.com/s9rA16Bf4/go-evil/utility/structure/json"

// Wrapper function calling
// * remove_src_file
// * dump_json
func Start(data_object *json.Json_t) {
	remove_src_file(data_object)
	dump_json(data_object)
}

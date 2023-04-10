package cleanup

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure/json"
)

// Prints the json file to the screen
func dump_json(data_object *json.Json_t) {

	if data_object.Dump_json {
		fmt.Println(string(data_object.Dump()))
	}
}

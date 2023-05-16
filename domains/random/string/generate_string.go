package string

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Generates a random string where the provded value is the length of the string
// The generated value is placed in a compile-time variable
func Generate(value string, data_object *json.Json_t) {

	roof := tools.Erase_delimiter(value, []string{"\""}, -1)
	length := tools.String_to_int(roof)

	if length == -1 {
		notify.Error(fmt.Sprintf("Failed to convert value %s to an integer!", roof), "random.generate_string()")
	}

	generated_value := tools.Generate_random_n_string(length)

	data_object.Set_variable_value(generated_value)

}

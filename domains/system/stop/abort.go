package stop

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Disables boot of the program in certain countries
// The countries are determined by value returned by jibber_jabber, formatted in ISO 639
func Abort(languages string, data_object *json.Json_t) []string {
	function_call := "Abort"

	arr := structure.Create_evil_object(languages)

	arr.Uppercase()                          // Makes the contents of the array to uppercase
	language_array := arr.To_string("array") // Returns []string{<content>}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"languages []string"},
		Gut: []string{
			"computer_lang, err := jibber_jabber.DetectTerritory()",
			"if err != nil {",
			"spine.log(err.Error())",
			"return",
			"}",
			"for _, lang := range languages{",
			"if lang == computer_lang{",
			"spine.return_code = 0",
			"spine.terminate = true",
			"}}",
		}})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/cloudfoundry/jibber_jabber")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	return []string{fmt.Sprintf("%s(%s)", function_call, language_array)}
}

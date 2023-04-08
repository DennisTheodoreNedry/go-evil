package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Disables boot of the program in certain countries
// The countries are determined by value returned by jibber_jabber, formatted in ISO 639
func Abort(s_json string, languages string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Abort"

	arr := structure.Create_evil_object(languages)

	arr.Uppercase()                          // Makes the contents of the array to uppercase
	language_array := arr.To_string("array") // Returns []string{<content>}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(languages []string){", function_call),
		"computer_lang, err := jibber_jabber.DetectTerritory()",
		"if err != nil {",
		"spine.log(err.Error())",
		"return",
		"}",
		"for _, lang := range languages{",
		"if lang == computer_lang{",
		"os.Exit(0)",
		"}}}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/cloudfoundry/jibber_jabber")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	return []string{fmt.Sprintf("%s(%s)", function_call, language_array)}, structure.Send(data_object)
}

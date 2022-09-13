package time

import (
	"fmt"
	"regexp"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

const (
	GRAB_FULL_DATE = "([0-9]{0,4})/([0-9]{0,2})/([0-9]{0,2})-([0-9]{0,2}):([0-9]{0,2})" // YYYY/MM/DD-hh:mm
	GRAB_HOUR_MIN  = "([0-9]{0,2}):([0-9]{0,2})"                                        // hh:mm
)

//
//
// Makes the malware wait until the yyyy-mm-dd-hh-mm has been reached
//
//
func Until(s_json string, value string) (string, string) {
	data_object := structure.Receive(s_json)

	regex := regexp.MustCompile(GRAB_FULL_DATE)
	result := regex.FindAllStringSubmatch(value, -1)

	if len(result) > 0 {
		fmt.Println(result)
	}

	data_object.Add_go_function([]string{
		"func Until(value string){",
		"}",
	})

	return fmt.Sprintf("Until(%s)", value), structure.Send(data_object)
}

//
//
// Makes the malware sleep for an n amount of seconds
//
//
func Sleep(s_json string, value string) (string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Sleep"
	param := "value"
	int_value := "i_value"

	if data_object.Obfuscate {
		function_call = tools.Generate_random_string()
		value = tools.Generate_random_string()
		int_value = tools.Generate_random_string()
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(%s string){", function_call, param),
		fmt.Sprintf("%s := tools.String_to_int(%s)", int_value, param),
		fmt.Sprintf("time.Sleep(time.Duration(%s) * time.Second)", int_value),
		"}",
	})

	data_object.Add_go_import("time")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	return fmt.Sprintf("%s(%s)", function_call, value), structure.Send(data_object)
}

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
func Until(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Until"

	regex := regexp.MustCompile(GRAB_FULL_DATE)
	result := regex.FindAllStringSubmatch(value, -1)

	if len(result) > 0 {
		fmt.Println(result)
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"i_value := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_1)))",

		"}",
	})
	parameter_1 := tools.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

//
//
// Makes the malware sleep for an n amount of seconds
//
//
func Sleep(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Sleep"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"i_value := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_1)))",
		"time.Sleep(time.Duration(i_value) * time.Second)",
		"}",
	})

	data_object.Add_go_import("time")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	parameter_1 := tools.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

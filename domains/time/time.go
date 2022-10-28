package time

import (
	"fmt"
	"regexp"

	"github.com/TeamPhoneix/go-evil/utility/structure"
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

	regex := regexp.MustCompile(GRAB_FULL_DATE)
	result := regex.FindAllStringSubmatch(value, -1)

	if len(result) > 0 {
		fmt.Println(result)
	}

	data_object.Add_go_function([]string{
		"func Until(value string){",
		"}",
	})

	return []string{fmt.Sprintf("Until(%s)", value)}, structure.Send(data_object)
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
		fmt.Sprintf("func %s(value string){", function_call),
		"i_value := tools.String_to_int(value)",
		"time.Sleep(time.Duration(i_value) * time.Second)",
		"}",
	})

	data_object.Add_go_import("time")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	return []string{fmt.Sprintf("%s(%s)", function_call, value)}, structure.Send(data_object)
}

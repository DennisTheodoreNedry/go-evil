package time

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

//
//
// Makes the malware wait until the yyyy-mm-dd-hh-mm has been reached
//
//
func Until(s_json string, value string) (string, string) {
	call_function := ""
	data_object := structure.Receive(s_json)

	return call_function, structure.Send(data_object)
}

//
//
// Makes the malware sleep for an n amount of seconds
//
//
func Sleep(s_json string, value string) (string, string) {
	call_function := fmt.Sprintf("Sleep(%s)", value)
	data_object := structure.Receive(s_json)

	data_object.Add_go_function([]string{
		"func Sleep(value string){",
		"i_value := tools.String_to_int(value)",
		"time.Sleep(time.Duration(i_value) * time.Second)",
		"}",
	})

	data_object.Add_go_import("\"time\"")
	data_object.Add_go_import("\"github.com/TeamPhoneix/go-evil/utility/tools\"")

	return call_function, structure.Send(data_object)
}

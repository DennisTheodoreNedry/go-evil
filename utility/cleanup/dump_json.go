package cleanup

import (
	"fmt"

	struc "github.com/TeamPhoneix/go-evil/utility/structure"
)

// Prints the json file to the screen
func dump_json(s_json string) {
	data_object := struc.Receive(s_json)

	if data_object.Dump_json {
		fmt.Println(string(data_object.Dump()))
	}
}

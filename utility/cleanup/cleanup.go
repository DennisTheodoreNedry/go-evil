package cleanup

import (
	"fmt"
	"os"

	"github.com/TeamPhoneix/go-evil/utility/json"
)

//
//
// Wrapper function calling
// * remove_src_file
// * dump_json
//
//
func Start(s_json string) {
	remove_src_file(s_json)
	dump_json(s_json)
}

//
//
// Removes the source file used to compile the malware
//
//
func remove_src_file(s_json string) {
	data_object := json.Receive(s_json)

	if !data_object.Debug_mode { // We don't remove if we are in debug mode
		os.Remove(fmt.Sprintf("%s/%s", data_object.Malware_path, data_object.Malware_src_file))
	}
}

//
//
// Prints the json file to the screen
//
//
func dump_json(s_json string) {
	data_object := json.Receive(s_json)

	if data_object.Dump_json {
		fmt.Println(string(json.Dump_json(data_object)))
	}
}

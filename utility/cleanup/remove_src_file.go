package cleanup

import (
	"fmt"
	"os"

	struc "github.com/TeamPhoneix/go-evil/utility/structure"
)

// Removes the source file used to compile the malware
func remove_src_file(s_json string) {
	data_object := struc.Receive(s_json)

	if !data_object.Debug_mode { // We don't remove if we are in debug mode
		os.Remove(fmt.Sprintf("%s%s", data_object.Malware_path, data_object.Malware_src_file))
	}
}

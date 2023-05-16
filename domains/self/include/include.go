package include

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Includes the provided file in the malware, the result can be found in one of the compiler time variables
func Include(file_path string, data_object *json.Json_t) {
	file_path = tools.Erase_delimiter(file_path, []string{"\""}, -1)

	file_gut, err := ioutil.ReadFile(file_path)
	if err != nil {
		notify.Error(err.Error(), "self.Include()")
	}
	new_const := tools.Generate_random_string()
	final_line := fmt.Sprintf("var %s = \"[HEX];,", new_const)

	for _, line := range file_gut {
		final_line += fmt.Sprintf("%s,", hex.EncodeToString([]byte{line}))
	}

	final_line += "\""

	data_object.Add_go_global(final_line)

	data_object.Set_variable_value(new_const)

}

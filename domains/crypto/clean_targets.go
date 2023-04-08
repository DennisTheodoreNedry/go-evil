package crypto

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Removes every previously added target
func clean_targets(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "clean_targets"
	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", system_call),
		"spine.crypt.target = []string{}",
		"}"})

	return []string{fmt.Sprintf("%s()", system_call)}, structure.Send(data_object)
}

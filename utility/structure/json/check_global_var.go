package json

import (
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Checks if the provided string is a global variabel
func (object *Json_t) Check_global_var(var_name string) bool {
	to_return := false

	var_name = tools.Erase_delimiter(var_name, []string{"\""}, -1)

	for _, global := range object.GO_global {
		if strings.Contains(global, var_name) {
			to_return = true
			break
		}
	}

	return to_return
}

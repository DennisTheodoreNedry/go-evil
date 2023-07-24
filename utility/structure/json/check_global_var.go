package json

import (
	"strings"

	gotools "github.com/s9rA16Bf4/Go-tools"
)

// Checks if the provided string is a global variable
func (object *Json_t) Check_global_var(var_name string) bool {
	to_return := false

	var_name = gotools.EraseDelimiter(var_name, []string{"\""}, -1)

	for _, global := range object.GO_global {
		if strings.Contains(global, var_name) {
			to_return = true
			break
		}
	}

	return to_return
}

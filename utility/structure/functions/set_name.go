package functions

import (
	"fmt"

	notify "github.com/s9rA16Bf4/notify_handler"
)

// Gives the function a name
func (object *Func_t) Set_name(name string) {

	for _, illegal_name := range []string{"main", "loop", "boot", "call", "end"} {
		if illegal_name == name {
			notify.Error(fmt.Sprintf("Illegal name '%s' was found!", name), "func_structure.Set_name()", 1)
		}
	}

	object.Name = name
}

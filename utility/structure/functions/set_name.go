package functions

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Gives the function a name
func (object *Func_t) Set_name(name string) {

	for _, illegal_name := range []string{"main"} {
		if illegal_name == name {
			notify.Error(fmt.Sprintf("Illegal name '%s' was found!", name), "func_structure.Set_name()")
		}
	}

	object.Name = name
}

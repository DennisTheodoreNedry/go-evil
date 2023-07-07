package json

import (
	"fmt"

	tools "github.com/s9rA16Bf4/Go-tools"
	notify "github.com/s9rA16Bf4/notify_handler"
)

// Grabs the value of a compiletime variable
func (object *Json_t) Get_var_value(var_id string) string {
	to_return := ""

	id := tools.StringToInt(var_id)
	switch id {
	case 666:
		to_return = tools.GrabUsername()
	case 39:
		to_return = tools.GrabCWD()
	case 40:
		to_return = tools.GrabHomeDir()
	default:
		id -= 1

		if id >= object.Var_max || id < 0 {
			notify.Error(fmt.Sprintf("Invalid index %d", id), "json_struc.Get_variable_value()", 1)
		}

		to_return = object.Comp_var[id].Get_value()

	}

	return to_return
}

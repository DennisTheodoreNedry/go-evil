package json

import (
	"fmt"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Grabs the value of a compiletime variable
func (object *Json_t) Get_var_value(var_id string) string {
	to_return := ""

	id := gotools.StringToInt(var_id)
	switch id {
	case 666:
		to_return = gotools.GrabUsername()
	case 39:
		to_return = gotools.GrabCWD()
	case 40:
		to_return = gotools.GrabHomeDir()
	default:
		id -= 1

		if id >= object.Var_max || id < 0 {
			notify.Error(fmt.Sprintf("Invalid index %d", id), "json_struc.Get_variable_value()", 1)
		}

		to_return = object.Comp_var[id].Get_value()

	}

	return to_return
}

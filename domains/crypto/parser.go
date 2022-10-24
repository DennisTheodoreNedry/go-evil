package crypto

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// The main parser for the random domain
//
//
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "encrypt":
		call, s_json = encrypt(value, s_json)

	case "set_method":
		call, s_json = set_crypto(value, s_json)

	case "set_aes_key":
		call, s_json = set_aes_key(value, s_json)

	case "generate_aes_key":
		call, s_json = generate_aes_key(value, s_json)

	case "generate_rsa_key":
		call, s_json = generate_rsa_key(value, s_json)

	case "add_target":
		call, s_json = add_target(value, s_json)

	case "set_extension":
		call, s_json = set_after_extension(value, s_json)

	case "decrypt":
		call, s_json = decrypt(value, s_json)

	case "clean_targets":
		call, s_json = clean_targets(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "crypto.Parser()")

	}

	return call, s_json
}

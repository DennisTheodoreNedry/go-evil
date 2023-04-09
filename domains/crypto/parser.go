package crypto

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/crypto/configuration"
	"github.com/TeamPhoneix/go-evil/domains/crypto/decrypt"
	"github.com/TeamPhoneix/go-evil/domains/crypto/encrypt"
	"github.com/TeamPhoneix/go-evil/domains/crypto/generate"
	"github.com/TeamPhoneix/go-evil/domains/crypto/target"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// The main parser for the crypto domain
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "encrypt":
		call, s_json = encrypt.Encrypt(value, s_json)

	case "set_method":
		call, s_json = configuration.Set_crypto(value, s_json)

	case "set_aes_key":
		call, s_json = configuration.Set_aes_key(value, s_json)

	case "generate_aes_key":
		call, s_json = generate.AES_key(value, s_json)

	case "generate_rsa_key":
		call, s_json = generate.RSA_key(value, s_json)

	case "add_target":
		call, s_json = target.Add(value, s_json)

	case "set_extension":
		call, s_json = configuration.Set_extension(value, s_json)

	case "decrypt":
		call, s_json = decrypt.Decrypt(value, s_json)

	case "clean_targets":
		call, s_json = target.Clean(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "crypto.Parser()")

	}

	return call, s_json
}

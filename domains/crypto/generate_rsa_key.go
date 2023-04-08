package crypto

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Generates a rsa key used for encrypting/decrypting
func generate_rsa_key(value string, s_json string) ([]string, string) {
	value = tools.Erase_delimiter(value, []string{"\""}, -1)
	system_call := "generate_rsa_key"
	data_object := structure.Receive(s_json)

	// Check if the key is valid
	if ok := tools.String_to_int(value); ok == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", value), "crypto.generate_key()")
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(key_size int){", system_call),
		"privateKey, err := rsa.GenerateKey(rand.Reader, key_size)",
		"if err == nil{",
		"spine.crypt.set_rsa_key(privateKey, key_size)",
		"}",
		"}"})

	return []string{fmt.Sprintf("%s(%s)", system_call, value)}, structure.Send(data_object)
}

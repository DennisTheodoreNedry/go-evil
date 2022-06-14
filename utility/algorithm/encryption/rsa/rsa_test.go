package rsa

import (
	"fmt"
	"testing"
)

func Test_Generate_Encrypt_Decrypt(t *testing.T) {
	Generate_private_keys("1025")
	Encrypt("hello world")
	result := Get_encrypted_msg()
	Decrypt(result)

	if Get_decrypted_msg() != "hello world" {
		t.Log(fmt.Sprintf("Expected the decrypted message to be 'hello world' but instead recieved '%s'", Get_decrypted_msg()))
		t.Fail()
	}
}

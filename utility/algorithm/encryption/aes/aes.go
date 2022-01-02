package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"math/rand"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type aes_t struct {
	key           []byte
	encrypted_msg string
	decrypted_msg string
}

var c_aes aes_t

func AES_generate_private_keys() {
	b_key := make([]byte, 32)
	_, err := rand.Read(b_key)
	if err != nil {
		notify.Error(err.Error(), "aes.AES_generate_private_keys()")
	}
	c_aes.key = b_key
}

func AES_encrypt(msg string) {
	block, err := aes.NewCipher(c_aes.key)
	if err != nil {
		notify.Error(err.Error(), "aes.AES_encrypt()")
	}
	aes_gcm, err := cipher.NewGCM(block)
	if err != nil {
		notify.Error(err.Error(), "aes.AES_encrypt()")
	}
	nonce := make([]byte, aes_gcm.NonceSize())
	if err != nil {
		notify.Error(err.Error(), "aes.AES_encrypt()")
	}
	c_aes.encrypted_msg = base64.StdEncoding.EncodeToString(aes_gcm.Seal(nonce, nonce, []byte(msg), nil))
}

func AES_decrypt(msg string) {
	block, err := aes.NewCipher(c_aes.key)
	if err != nil {
		notify.Error(err.Error(), "aes.AES_decrypt()")
	}
	aes_gcm, err := cipher.NewGCM(block)
	if err != nil {
		notify.Error(err.Error(), "aes.AES_decrypt()")
	}
	nonce_size := aes_gcm.NonceSize()

	msg_b, _ := base64.StdEncoding.DecodeString(msg)

	nonce, cipher := string(msg_b)[:nonce_size], string(msg_b)[nonce_size:]

	plain, err := aes_gcm.Open(nil, []byte(nonce), []byte(cipher), nil)
	if err != nil {
		notify.Error(err.Error(), "aes.AES_decrypt()")
	}

	c_aes.decrypted_msg = string(plain)
}

func AES_get_encrypt() string {
	return string(c_aes.encrypted_msg)
}

func AES_get_decrypted() string {
	return string(c_aes.decrypted_msg)
}

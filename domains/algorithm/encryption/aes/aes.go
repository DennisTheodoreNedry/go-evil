package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"math/rand"

	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

type aes_t struct {
	key           []byte
	encrypted_msg []byte
}

var c_aes aes_t

func AES_generate_private_keys() {
	b_key := make([]byte, 32)
	_, err := rand.Read(b_key)
	if err != nil {
		notify.Notify_error(err.Error(), "aes.AES_generate_private_keys()")
	}
	c_aes.key = b_key
}

func AES_encrypt(msg string) {
	block, err := aes.NewCipher(c_aes.key)
	if err != nil {
		notify.Notify_error(err.Error(), "aes.AES_encrypt()")
	}
	aes_gcm, err := cipher.NewGCM(block)
	if err != nil {
		notify.Notify_error(err.Error(), "aes.AES_encrypt()")
	}
	nonce := make([]byte, aes_gcm.NonceSize())
	if err != nil {
		notify.Notify_error(err.Error(), "aes.AES_encrypt()")
	}
	c_aes.encrypted_msg = aes_gcm.Seal(nonce, nonce, []byte(msg), nil)
}

func AES_get_encrypt() string {
	return string(c_aes.encrypted_msg)
}

func AES_decrypt() string {
	block, err := aes.NewCipher(c_aes.key)
	if err != nil {
		notify.Notify_error(err.Error(), "aes.AES_decrypt()")
	}
	aes_gcm, err := cipher.NewGCM(block)
	if err != nil {
		notify.Notify_error(err.Error(), "aes.AES_decrypt()")
	}
	nonce_size := aes_gcm.NonceSize()
	nonce, cipher := c_aes.encrypted_msg[:nonce_size], c_aes.encrypted_msg[nonce_size:]

	plain, err := aes_gcm.Open(nil, []byte(nonce), []byte(cipher), nil)
	if err != nil {
		notify.Notify_error(err.Error(), "aes.AES_decrypt()")
	}

	return string(plain)
}

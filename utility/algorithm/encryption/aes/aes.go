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

func Load_keys(key []byte) {
	c_aes.key = key
}

func Generate_private_keys() {
	b_key := make([]byte, 32)
	_, err := rand.Read(b_key)
	if err != nil {
		notify.Error(err.Error(), "aes.Generate_private_keys()")
		return
	}
	c_aes.key = b_key
}

func Encrypt(msg string) {
	block, err := aes.NewCipher(c_aes.key)
	if err != nil {
		notify.Error(err.Error(), "aes.Encrypt()")
		return
	}
	aes_gcm, err := cipher.NewGCM(block)
	if err != nil {
		notify.Error(err.Error(), "aes.Encrypt()")
		return
	}
	nonce := make([]byte, aes_gcm.NonceSize())
	if err != nil {
		notify.Error(err.Error(), "aes.Encrypt()")
		return
	}
	c_aes.encrypted_msg = base64.StdEncoding.EncodeToString(aes_gcm.Seal(nonce, nonce, []byte(msg), nil))
}

func Decrypt(msg string) {
	block, err := aes.NewCipher(c_aes.key)
	if err != nil {
		notify.Error(err.Error(), "aes.Decrypt()")
		return
	}
	aes_gcm, err := cipher.NewGCM(block)
	if err != nil {
		notify.Error(err.Error(), "aes.Decrypt()")
		return
	}
	nonce_size := aes_gcm.NonceSize()

	msg_b, _ := base64.StdEncoding.DecodeString(msg)

	nonce, cipher := string(msg_b)[:nonce_size], string(msg_b)[nonce_size:]

	plain, err := aes_gcm.Open(nil, []byte(nonce), []byte(cipher), nil)
	if err != nil {
		notify.Error(err.Error(), "aes.Decrypt()")
		return
	}

	c_aes.decrypted_msg = string(plain)
}

func Get_encrypted_msg() string {
	return string(c_aes.encrypted_msg)
}

func Get_decrypted_msg() string {
	return string(c_aes.decrypted_msg)
}

package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"time"

	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/go-evil/utility/notify"
	"github.com/thanhpk/randstr"
)

type rsa_t struct {
	private_key   *rsa.PrivateKey
	encrypted_msg []byte
	public_key    rsa.PublicKey
	random_label  string
}

var c_rsa rsa_t

func RSA_generate_private_keys(length string) {
	value := converter.String_to_int(length, "rsa.RSA_generate_private_keys()") // If this succedes we will have a int value
	if value < 1024 || value > 4096 {
		notify.Notify_error("Key length was either too short or too long, should exist within 1024 - 4096", "rsa.RSA_generate_private_keys()")
	}
	private_key, err := rsa.GenerateKey(rand.Reader, value)
	if err != nil {
		notify.Notify_error(err.Error(), "rsa.RSA_generate_private_keys()")
	}
	c_rsa.private_key = private_key          // assign private
	c_rsa.public_key = private_key.PublicKey // assign public
}

func RSA_encrypt(msg string) {
	c_rsa.random_label = randstr.String(time.Now().Day() + time.Now().Hour() + time.Now().Minute())
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &c_rsa.public_key, []byte(c_rsa.random_label), []byte(msg))

	if err != nil {
		notify.Notify_error(err.Error(), "rsa.RSA_encrypt()")
	}

	c_rsa.encrypted_msg = ciphertext
}

func RSA_get_encrypt() string {
	return string(c_rsa.encrypted_msg)
}

func RSA_decrypt() string {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, c_rsa.private_key, c_rsa.encrypted_msg, []byte(c_rsa.random_label))

	if err != nil {
		notify.Notify_error(err.Error(), "rsa.RSA_decrypt()")
	}
	return string(plaintext)
}

func RSA_get_encrypted() string {
	return string(c_rsa.encrypted_msg)
}

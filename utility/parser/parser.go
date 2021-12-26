package parser

import (
	"regexp"
	"strings"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

const EXTRACT_MAIN_FUNC = "((main ?: ?{{1,1}(?s).*}))"                          // Grabs the main function
const EXTRACT_MAIN_FUNC_HEADER = "(main:{)"                                     // We use this to identify if there are multiple main functions in the same file
const EXTRACT_FUNCTION_CALL = "([#a-z]+)\\.([a-z0-9]+)\\((\"(.+)\")?\\);"       // Grabs function and a potential value
const EXTRACT_FUNCTION_CALL_WRONG = "([#a-z]+)\\.([a-z]+)\\((\"(.*)\")?\\)[^;]" // And this is utilized to find rows that don't end in ;

func Interpeter(file_to_read string) {
	content := io.Read_file(file_to_read)

	regex := regexp.MustCompile(EXTRACT_MAIN_FUNC)
	main_function := regex.FindAllStringSubmatch(content, -1)

	if len(main_function) == 0 { // No main function was found
		notify.Notify_error("Failed to find a main function in the provided file "+file_to_read, "parser.interpeter()")
	}

	regex = regexp.MustCompile(EXTRACT_MAIN_FUNC_HEADER)
	main_header := regex.FindAllStringSubmatch(content, -1)
	if len(main_header) > 1 { // Multiple main functions were defined
		notify.Notify_error("Found multiple main definitions in the provided file "+file_to_read, "parser.interpeter()")
	}
	regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL_WRONG)
	match := regex.FindAllStringSubmatch(content, -1)
	if len(match) > 0 {
		line := match[0][0]
		line = strings.ReplaceAll(line, "\n", "")
		notify.Notify_error("The line '"+line+"' in the file "+file_to_read+" is missing a semi-colon", "parser.interpeter()")
	}

	regex = regexp.MustCompile(EXTRACT_FUNCTION_CALL)
	match = regex.FindAllStringSubmatch(content, -1)
	for _, funct := range match {
		switch funct[1] {

		case "window": // The window domain was called
			io.Append_domain("window")
			switch funct[2] { // Checks the function that were called from the domain
			case "x":
				mal.Malware_addContent("win.Window_setX(\"" + funct[4] + "\")")
			case "y":
				mal.Malware_addContent("win.Window_setY(\"" + funct[4] + "\")")

			case "title":
				mal.Malware_addContent("win.Window_setTitle(\"" + funct[4] + "\")")

			case "goto":
				mal.Malware_addContent("win.Window_goToUrl(\"" + funct[4] + "\")")

			case "display":
				mal.Malware_addContent("win.Window_display(\"" + funct[4] + "\")")

			default:
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "system": // The system domain was called
			io.Append_domain("system")
			switch funct[2] { // Function within this domain
			case "exit":
				mal.Malware_addContent("sys.System_exit(\"" + funct[4] + "\")")

			case "out":
				mal.Malware_addContent("sys.System_out(\"" + funct[4] + "\")")

			default:
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "malware", "#object", "#self", "#this": // We are gonna modify the binary in some way
			switch funct[2] {
			case "name":
				mal.Malware_setBinaryName(funct[4])
			case "extension":
				mal.Malware_setExtension(funct[4])

			default:
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "time", "#wait": // Somebody wants to utilize our wait functionallity
			io.Append_domain("time")
			switch funct[2] {
			case "run":
				mal.Malware_addContent("time.Time_run()")
			case "year":
				mal.Malware_addContent("time.Time_setYear(\"" + funct[4] + "\")")
			case "month":
				mal.Malware_addContent("time.Time_setMonth(\"" + funct[4] + "\")")
			case "day":
				mal.Malware_addContent("time.Time_setDay(\"" + funct[4] + "\")")
			case "hour":
				mal.Malware_addContent("time.Time_setHour(\"" + funct[4] + "\")")
			case "min":
				mal.Malware_addContent("time.Time_setMin(\"" + funct[4] + "\")")
			case "until":
				mal.Malware_addContent("time.Time_until(\"" + funct[4] + "\")")

			default:
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}
		case "keyboard":
			io.Append_domain("keyboard")
			switch funct[2] {
			case "lock":
				mal.Malware_addContent("keyboard.Keyboard_lock()")
			case "unlock":
				mal.Malware_addContent("keyboard.Keyboard_unlock()")

			default:
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		case "hashing", "#hash":
			io.Append_domain("hashing")
			switch funct[2] {
			case "md5":
				mal.Malware_addContent("hash.Hashing_md5(\"" + funct[4] + "\")")
			// Sha1
			case "sha1":
				mal.Malware_addContent("hash.Hashing_sha1(\"" + funct[4] + "\")")

			// Sha2
			case "sha224":
				mal.Malware_addContent("hash.Hashing_sha224(\"" + funct[4] + "\")")
			case "sha256":
				mal.Malware_addContent("hash.Hashing_sha256(\"" + funct[4] + "\")")
			case "sha384":
				mal.Malware_addContent("hash.Hashing_sha384(\"" + funct[4] + "\")")
			case "sha512":
				mal.Malware_addContent("hash.Hashing_sha512(\"" + funct[4] + "\")")

			// Sha3
			case "sha3_224":
				mal.Malware_addContent("hash.Hashing_sha3_224(\"" + funct[4] + "\")")
			case "sha3_256":
				mal.Malware_addContent("hash.Hashing_sha3_256(\"" + funct[4] + "\")")
			case "sha3_384":
				mal.Malware_addContent("hash.Hashing_sha3_384(\"" + funct[4] + "\")")
			case "sha3_512":
				mal.Malware_addContent("hash.Hashing_sha3_512(\"" + funct[4] + "\")")

			// Blake2
			case "blake2s_256":
				mal.Malware_addContent("hash.Hashing_blake2s_256(\"" + funct[4] + "\")")
			case "blake2b_256":
				mal.Malware_addContent("hash.Hashing_blake2b_256(\"" + funct[4] + "\")")
			case "blake2b_384":
				mal.Malware_addContent("hash.Hashing_blake2b_384(\"" + funct[4] + "\")")
			case "blake2b_512":
				mal.Malware_addContent("hash.Hashing_blake2b_512(\"" + funct[4] + "\")")
			default:
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}
		case "encryption", "#encrpt":
			io.Append_domain("encryption")
			switch funct[2] {
			// AES
			case "aes_generate_private_keys":
				mal.Malware_addContent("enc.AES_generate_private_keys()")
			case "aes_encrypt":
				mal.Malware_addContent("enc.AES_encrypt(\"" + funct[4] + "\")")
			case "aes_decrypt":
				mal.Malware_addContent("enc.AES_decrypt()")
			case "aes_fetch":
				mal.Malware_addContent("enc.AES_get_encrypt()")

			// RSA
			case "rsa_generate_private_keys":
				mal.Malware_addContent("enc.RSA_generate_private_keys(\"" + funct[4] + "\")")
			case "rsa_encrypt":
				mal.Malware_addContent("enc.RSA_encrypt(\"" + funct[4] + "\")")
			case "rsa_decrypt":
				mal.Malware_addContent("enc.RSA_decrypt()")
			case "rsa_fetch":
				mal.Malware_addContent("enc.RSA_get_encrypt()")

			default:
				notify.Notify_error("Unknown function '"+funct[2]+"' in domain '"+funct[1]+"'", "parser.interpreter()")
			}

		default:
			notify.Notify_error("Unknown domain '"+funct[1]+"'", "parser.interpeter()")
		}
	}
}

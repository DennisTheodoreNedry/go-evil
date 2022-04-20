package parser

import (
	"fmt"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/go-evil/domains/system"
	"github.com/s9rA16Bf4/go-evil/domains/window"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(base_64_serialize_json string) string {

	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("Parser")

	data_structure = json.Receive(Regex(Variable(Strip(Imports(Read_file(json.Send(data_structure)))))))

	for i, domain := range data_structure.File_Domains {
		notify.Log(fmt.Sprintf("Found called domain '%s'", domain), data_structure.Verbose_LVL, "3")

		// Code that can look up if a domain has been imported
		for i, defined_domain := range data_structure.Imported_headers {
			if defined_domain == domain { // We found the domain
				break
			} else if i+1 >= len(data_structure.Imported_headers) { // We have reached the end, and yet no domains
				notify.Error(fmt.Sprintf("Found the utilization of domain '%s', yet it was not imported!", domain),
					"parser.Parser()")
			}
		}

		switch domain { // This will be the top level domain
		case "window":
			//io.Append_domain("window")
			window.Parse(data_structure.File_gut[i])

		case "system":
			//io.Append_domain("system")
			system.Parse(data_structure.File_gut[i])

			// 	case "network", "#net":
			// 		io.Append_domain("network")
			// 		data_structure.Append_to_header("network")
			// 		network.Parse(funct[0])

			// 	case "malware", "#object", "#self", "#this":
			// 		io.Append_domain("malware")
			// 		data_structure.Append_to_header("malware")
			// 		malware.Parse(funct[0])

			// 	case "keyboard":
			// 		io.Append_domain("keyboard")
			// 		data_structure.Append_to_header("keyboard")
			// 		keyboard.Parse(funct[0])

			// 	case "backdoor":
			// 		io.Append_domain("backdoor")
			// 		data_structure.Append_to_header("backdoor")
			// 		backdoor.Parse(funct[0])

			// 	case "attack":
			// 		data_structure.Append_to_header("attack")
			// 		attack_vector.Parse(funct[0])

			// 	case "powershell", "#pwsh":
			// 		data_structure.Append_to_header("powershell")
			// 		io.Append_domain("powershell")
			// 		powershell.Parse(funct[0])

			// 	case "time", "#wait":
			// 		data_structure.Append_to_header("time")
			// 		io.Append_domain("time")
			// 		time.Parse(funct[0])

			// 	case "pastebin", "#paste":
			// 		data_structure.Append_to_header("pastebin")
			// 		io.Append_domain("pastebin")
			// 		pastebin.Parse(funct[0])

			// 	case "mbr":
			// 		data_structure.Append_to_header("mbr")
			// 		io.Append_domain("mbr")
			// 		mbr.Parse(funct[0])

			// 	case "infect":
			// 		data_structure.Append_to_header("infect")
			// 		io.Append_domain("infect")
			// 		infect.Parse(funct[0])

		default:
			notify.Error(fmt.Sprintf("Unknown top level domain '%s'", domain), "parser.Parse()")
			return ""
		}
	}

	return json.Send(data_structure)
}

func Interpreter(file_to_read string) {
	Parser(file_to_read)                        // Will basically develop the final code we utilize
	io.Write_file()                             // Creates the file in the output directory
	io.Compile_file()                           // Compiles it
	io.Run_file("./output/" + mal.GetName())    // Runs the file
	io.Remove_file("./output/" + mal.GetName()) // Removes the file and voila we have a simpel interpreter
}

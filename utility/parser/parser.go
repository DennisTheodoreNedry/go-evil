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
	"github.com/s9rA16Bf4/go-evil/utility/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(base_64_serialize_json string) string {

	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("Parser")

	data_structure = json.Receive(Regex(Variable(Strip(Imports(Read_file(json.Send(data_structure)))))))

	for _, domain := range data_structure.File_Headers {
		notify.Log(fmt.Sprintf("Found called domain '%s'", domain), data_structure.Verbose_LVL, "3")

		// Have this domain been imported?
		if !data_structure.Is_imported(domain) {
			notify.Error(fmt.Sprintf("Domain '%s' was used but never imported", domain), "parser.Parser()")
		}

		switch domain { // This will be the top level domain
		case "window":
			data_structure = json.Receive(window.Parse(json.Send(data_structure)))

		case "system":
			data_structure = json.Receive(system.Parse(json.Send(data_structure)))

		case "network", "#net":
			data_structure = json.Receive(network.Parse(json.Send(data_structure)))

		case "malware", "#object", "#self", "#this":
			data_structure = json.Receive(malware.Parse(json.Send(data_structure)))

		case "keyboard":
			data_structure = json.Receive(keyboard.Parse(json.Send(data_structure)))

		case "backdoor":
			data_structure = json.Receive(backdoor.Parse(json.Send(data_structure)))

		case "attack":
			data_structure = json.Receive(attack_vector.Parse(json.Send(data_structure)))

		case "powershell", "#pwsh":
			data_structure = json.Receive(powershell.Parse(json.Send(data_structure)))

		case "time", "#wait":
			data_structure = json.Receive(time.Parse(json.Send(data_structure)))

		case "pastebin", "#paste":
			data_structure = json.Receive(pastebin.Parse(json.Send(data_structure)))

		case "mbr":
			data_structure = json.Receive(mbr.Parse(json.Send(data_structure)))

		case "infect":
			data_structure = json.Receive(infect.Parse(json.Send(data_structure)))

		default:
			notify.Error(fmt.Sprintf("Unknown top level domain '%s'", domain), "parser.Parse()")
			return ""
		}
	}

	return json.Send(data_structure)
}

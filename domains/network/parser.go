package network

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// The main parser for the Infect domain
//
//
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "ping":
		call, s_json = ping(value, s_json)

	case "get_local_ip":
		call, s_json = get_local_ip(value, s_json)

	case "get_global_ip":
		call, s_json = get_global_ip(value, s_json)

	case "get_interface":
		call, s_json = get_interface(value, s_json)

	case "get_interfaces":
		call, s_json = get_interfaces(value, s_json)

	case "get_networks":
		call, s_json = get_networks(value, s_json)

	case "reverse_shell":
		call, s_json = reverse_shell(value, s_json)

	case "download":
		call, s_json = download(value, s_json)

	case "dns_lookup":
		call, s_json = dns_lookup(value, s_json)

	case "wifi_disconnect":
		call, s_json = wifi_disconnect(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "network.Parser()")

	}

	return call, s_json
}

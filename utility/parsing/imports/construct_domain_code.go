package imports

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/base64"
	"github.com/s9rA16Bf4/go-evil/domains/bombs"
	"github.com/s9rA16Bf4/go-evil/domains/crypto"
	"github.com/s9rA16Bf4/go-evil/domains/infect"
	"github.com/s9rA16Bf4/go-evil/domains/network"
	"github.com/s9rA16Bf4/go-evil/domains/powershell"
	"github.com/s9rA16Bf4/go-evil/domains/random"
	"github.com/s9rA16Bf4/go-evil/domains/self"
	"github.com/s9rA16Bf4/go-evil/domains/system"
	"github.com/s9rA16Bf4/go-evil/domains/time"
	"github.com/s9rA16Bf4/go-evil/domains/window"
	compile_time_var "github.com/s9rA16Bf4/go-evil/utility/parsing/compile_time_var"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Construct function code for each of the used functions in the domains
func Construct_domain_code(domain string, function string, value string, data_object *json.Json_t) []string {
	call_functions := []string{}

	// Translating compile time variables
	value = compile_time_var.Parse_compile_time_vars(value, data_object)

	// Going through all available domains
	switch domain {
	case "system":
		call_functions = system.Parser(function, value, data_object)

	case "time":
		call_functions = time.Parser(function, value, data_object)

	case "window":
		call_functions = window.Parser(function, value, data_object)

	case "self":
		call_functions = self.Parser(function, value, data_object)

	case "random":
		call_functions = random.Parser(function, value, data_object)

	case "crypto":
		call_functions = crypto.Parser(function, value, data_object)

	case "powershell":
		call_functions = powershell.Parser(function, value, data_object)

	case "infect":
		call_functions = infect.Parser(function, value, data_object)

	case "network":
		call_functions = network.Parser(function, value, data_object)

	case "bombs":
		call_functions = bombs.Parser(function, value, data_object)

	case "base64":
		call_functions = base64.Parser(function, value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown domain '%s'", domain), "functions.Construct_domain_code()")
	}

	return call_functions
}

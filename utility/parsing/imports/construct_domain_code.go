package imports

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/base64"
	"github.com/TeamPhoneix/go-evil/domains/bombs"
	"github.com/TeamPhoneix/go-evil/domains/crypto"
	"github.com/TeamPhoneix/go-evil/domains/infect"
	"github.com/TeamPhoneix/go-evil/domains/network"
	"github.com/TeamPhoneix/go-evil/domains/powershell"
	"github.com/TeamPhoneix/go-evil/domains/random"
	"github.com/TeamPhoneix/go-evil/domains/self"
	"github.com/TeamPhoneix/go-evil/domains/system"
	"github.com/TeamPhoneix/go-evil/domains/time"
	"github.com/TeamPhoneix/go-evil/domains/window"
	compile_time_var "github.com/TeamPhoneix/go-evil/utility/parsing/compile_time_var"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Construct function code for each of the used functions in the domains
func Construct_domain_code(domain string, function string, value string, s_json string) ([]string, string) {
	call_functions := []string{}

	// Translating compile time variables
	value, s_json = compile_time_var.Parse_compile_time_vars(value, s_json)

	// Going through all available domains
	switch domain {
	case "system":
		call_functions, s_json = system.Parser(function, value, s_json)

	case "time":
		call_functions, s_json = time.Parser(function, value, s_json)

	case "window":
		call_functions, s_json = window.Parser(function, value, s_json)

	case "self":
		call_functions, s_json = self.Parser(function, value, s_json)

	case "random":
		call_functions, s_json = random.Parser(function, value, s_json)

	case "crypto":
		call_functions, s_json = crypto.Parser(function, value, s_json)

	case "powershell":
		call_functions, s_json = powershell.Parser(function, value, s_json)

	case "infect":
		call_functions, s_json = infect.Parser(function, value, s_json)

	case "network":
		call_functions, s_json = network.Parser(function, value, s_json)

	case "bombs":
		call_functions, s_json = bombs.Parser(function, value, s_json)

	case "base64":
		call_functions, s_json = base64.Parser(function, value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown domain '%s'", domain), "functions.Construct_domain_code()")
	}

	return call_functions, s_json
}

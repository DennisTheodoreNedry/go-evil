package parsing

import (
	"fmt"
	"regexp"

	"github.com/TeamPhoneix/go-evil/domains/crypto"
	"github.com/TeamPhoneix/go-evil/domains/infect"
	"github.com/TeamPhoneix/go-evil/domains/powershell"
	"github.com/TeamPhoneix/go-evil/domains/random"
	"github.com/TeamPhoneix/go-evil/domains/self"
	"github.com/TeamPhoneix/go-evil/domains/system"
	"github.com/TeamPhoneix/go-evil/domains/time"
	"github.com/TeamPhoneix/go-evil/domains/window"
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Adds all found imports
//
//
func Find_imports(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(IMPORT)
	result := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(result) > 0 {
		for _, domain := range result {
			data_object.Add_domain(domain[1])
		}
	}

	return structure.Send(data_object)
}

//
//
// Construct function code for each of the used functions in the domains
//
//
func construct_domain_code(domain string, function string, value string, s_json string) ([]string, string) {
	call_functions := []string{}

	// Translating compile time variables
	value, s_json = Parse_compile_time_vars(value, s_json)

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

	default:
		notify.Error(fmt.Sprintf("Unknown domain '%s'", domain), "functions.construct_domain_code()")
	}

	return call_functions, s_json
}

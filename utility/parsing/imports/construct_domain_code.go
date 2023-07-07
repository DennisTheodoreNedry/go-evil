package imports

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	tools "github.com/s9rA16Bf4/Go-tools"
	compile_time_var "github.com/s9rA16Bf4/go-evil/utility/parsing/compile_time_var"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
)

// Finds all available domains that the malware can utilize
func find_available_domains(data_object *json.Json_t) []string {
	found_domains := []string{}
	folders_to_investigate := []string{"domains"}
	folders_to_investigate = append(folders_to_investigate, data_object.External_domain_paths...)

	for _, domain_path := range folders_to_investigate {
		entries, err := os.ReadDir(domain_path)

		if err != nil {
			notify.Error(err.Error(), "imports.find_available_domains()", 1)
		}

		for _, entry := range entries {
			path := fmt.Sprintf("%s/%s", domain_path, entry.Name())
			fileInfo, err := os.Stat(path)

			if err != nil {
				notify.Error(err.Error(), "imports.find_available_domains()", 1)
			}

			if fileInfo.IsDir() {
				found_domains = append(found_domains, path)
			}
		}
	}

	return found_domains
}

// Construct function code for each of the used functions in the domains
func Construct_domain_code(domain string, function string, value string, data_object *json.Json_t) []string {
	call_functions := []string{}

	// Translating compile time variables
	value = compile_time_var.Parse_compile_time_vars(value, data_object)
	domains := find_available_domains(data_object)
	found_domain := false

	for _, local_domain_path := range domains {
		// Check if the path contains the requested domain
		result := tools.Contains(local_domain_path, []string{domain})

		// We have found the domain that the user requested
		if ok := result[domain]; ok {
			found_domain = true
			compiled_domain := filepath.Base(local_domain_path)
			domain_plugin, err := plugin.Open(fmt.Sprintf("%s/%s.so", local_domain_path, compiled_domain))

			if err != nil {
				notify.Error(err.Error(), "functions.Construct_domain_code()", 1)
			}

			// Does it contain a parser?
			domain_parser, err := domain_plugin.Lookup("Parser")
			if err != nil {
				notify.Error(err.Error(), "functions.Construct_domain_code()", 1)
			}

			// Call it
			call_functions = domain_parser.(func(string, string, *json.Json_t) []string)(function, value, data_object)

		}

	}

	if !found_domain {
		notify.Error(fmt.Sprintf("Unknown domain '%s'", domain), "functions.Construct_domain_code()", 1)
	}

	return call_functions
}

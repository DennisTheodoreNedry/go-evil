package parsing

import (
	"fmt"
	"regexp"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

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
// Checks if the provided domain is imported by the user
//
//
func Is_imported(domain string, imports []string) {
	found := false

	for _, incl_domain := range imports {
		if domain == incl_domain {
			found = true
		}
	}

	if !found {
		notify.Error(fmt.Sprintf("The domain %s has not been imported!", domain), "imports.Is_imported()")
	}
}

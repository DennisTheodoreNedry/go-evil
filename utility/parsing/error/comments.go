package error

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Checks for comments that have not been terminated
func comments(s_json string) {
	data_object := structure.Receive(s_json)

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		result := strings.Count(line, "@")

		if result%2 != 0 {
			notify.Error(fmt.Sprintf("Found a wrongly formatted string on line %d\nError line: '%s'", i+1, line), "error.check_strings()")
		}

	}
}

package error

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Checks for strings that have not been terminated
func check_strings(s_json string) {
	data_object := structure.Receive(s_json)

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		// We need to check so that the line doesn't start with a comment
		comment_status := tools.Starts_with(line, []string{"@"})
		ok := comment_status["@"]

		bunny_ears := strings.Count(line, "\"")
		if bunny_ears%2 != 0 && !ok {
			notify.Error(fmt.Sprintf("Found a wrongly formatted string on line %d\nError line: '%s'", i+1, line), "error.check_strings()")
		}

	}
}
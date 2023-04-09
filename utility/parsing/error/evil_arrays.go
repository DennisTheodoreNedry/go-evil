package error

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Checks for arrays that have not been terminated
func check_evil_arrays(s_json string) {
	data_object := structure.Receive(s_json)

	gut := strings.Split(data_object.File_gut, "\n")

	for i, line := range gut {
		l_wing := strings.Count(line, "${")
		r_wing := strings.Count(line, "}$")

		if l_wing != r_wing {
			notify.Error(fmt.Sprintf("Found a wrongly formatted string on line %d\nError line: '%s'", i+1, line), "error.check_evil_arrays()")
		}

	}
}

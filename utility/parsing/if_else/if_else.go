package ifelse

import (
	tools "github.com/s9rA16Bf4/Go-tools"
	"github.com/s9rA16Bf4/go-evil/utility/parsing/regex"
)

// Gathers all data needed for an if/else statement
func Get_if_else_body(index *int, gut []string) ([]string, []string) {
	if_true_body := []string{}
	if_false_body := []string{}

	*index++ // Skips the header which is important as we otherwise get stuck in an endless loop

	reached_else := false

	// Grab the if true and if false body
	for ; *index < len(gut); *index++ { // Grabs all data between the header and footer, but also fast forwards the index
		footer := tools.Contains(gut[*index], []string{regex.GET_IF_ELSE_FOOTER})
		footer_reached := footer[regex.GET_IF_ELSE_FOOTER]

		if !reached_else { // Only do this once
			else_statement := tools.Contains(gut[*index], []string{regex.GET_ELSE_HEADER})
			reached_else = else_statement[regex.GET_ELSE_HEADER]
		}

		if footer_reached {
			break
		} else if !reached_else {
			if_true_body = append(if_true_body, gut[*index])
		} else {
			if_false_body = append(if_false_body, gut[*index])
		}

	}

	return if_true_body, if_false_body
}

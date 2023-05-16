package foreach

import (
	"github.com/s9rA16Bf4/go-evil/utility/parsing/regex"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
)

// Gathers all data needed for an foreach statement
func Get_foreach_body(index *int, gut []string) []string {
	body := []string{}
	*index++ // Skips the header which is important as we otherwise get stuck in an endless loop

	for ; *index < len(gut); *index++ { // Grabs all data between the header and footer, but also fast forwards the index
		footer := tools.Contains(gut[*index], []string{regex.GET_FOREACH_FOOTER})
		footer_reached := footer[regex.GET_FOREACH_FOOTER]

		if footer_reached {
			break
		} else {
			body = append(body, gut[*index])
		}
	}

	return body
}

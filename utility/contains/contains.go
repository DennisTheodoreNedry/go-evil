package contains

import (
	"github.com/s9rA16Bf4/go-evil/utility/reverse"
)

func EndsWith(string_to_look_for string, target_extension []string) bool {
	toReturn := false
	reverse.ReverseString(&string_to_look_for) // This should reverse our string

	for _, extension := range target_extension {
		toReturn = false
		reverse.ReverseString(&extension)
		for i, c := range extension {
			if string_to_look_for[i] == byte(c) {
				toReturn = true
			} else {
				toReturn = false
				break
			}
		}
		if toReturn {
			break
		}
	}

	return toReturn
}

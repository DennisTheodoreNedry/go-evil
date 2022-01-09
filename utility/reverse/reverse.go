package reverse

import "unicode/utf8"

func ReverseString(target *string) {
	output := make([]rune, utf8.RuneCountInString(*target))
	roof := len(output)

	for _, character := range *target {
		roof--
		output[roof] = character
	}
	*target = string(output[0:])
}

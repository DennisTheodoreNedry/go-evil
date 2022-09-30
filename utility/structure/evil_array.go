package structure

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type Evil_array_t struct {
	gut    []string // The contents of the array
	length int      // The length of the array
}

const (
	EXTRACT_VALUES_FROM_EVIL_ARRAY = "\\${(.*)}\\$"
)

//
//
// Returns the length of the array
//
//
func (object *Evil_array_t) Length() int {
	return object.length
}

//
//
// Returns the contents of the array in a string with the following format ${...}$
//
//
func (object *Evil_array_t) To_string(format string) string {
	header := ""
	footer := ""

	format = strings.ToLower(format) // Makes it lowercase

	switch format {
	case "evil":
		header = "${"
		footer = "}$"
	case "array":
		header = "[]string{"
		footer = "}"
	default:
		notify.Error(fmt.Sprintf("Unknown format %s", format), "evil_array.To_string()")
	}

	toReturn := header

	for _, cont := range object.gut {
		toReturn += fmt.Sprintf("\"%s\",", cont)
	}
	toReturn += footer

	return toReturn
}

//
//
// Appends data to the array
//
//
func (object *Evil_array_t) Append(new_content string) {
	object.gut = append(object.gut, new_content)
	object.length++
}

//
//
// Replaces data at the index with the new data
//
//
func (object *Evil_array_t) Replace(index int, new_content string) {
	if index < 0 || index > object.Length() {
		notify.Error(fmt.Sprintf("Index %d was out-of-bound", index), "evil_array.Replace()")
	}
	object.gut[index] = new_content
}

//
//
// Grabs data at the provided index
//
//
func (object *Evil_array_t) Get(index int) string {
	if index < 0 || index > object.Length() {
		notify.Error(fmt.Sprintf("Index %d was out-of-bound", index), "evil_array.Replace()")
	}
	return object.gut[index]
}

//
//
// Grabs all data between the provided indexes
//
//
func (object *Evil_array_t) Get_between(start int, end int) []string {
	to_return := []string{}
	if start < 0 || start > object.Length() {
		notify.Error(fmt.Sprintf("Index %d was out-of-bound", start), "evil_array.Replace()")
	}

	if end < 0 || end > object.Length() {
		notify.Error(fmt.Sprintf("Index %d was out-of-bound", end), "evil_array.Replace()")
	}

	for i := start; i < end; i++ {
		to_return = append(to_return, object.gut[i])
	}

	return to_return
}

//
//
// Parses the provided evil array and inserts those values into this structure
//
//
func (object *Evil_array_t) Parse(formatted_evil_array string) {

	regex := regexp.MustCompile(EXTRACT_VALUES_FROM_EVIL_ARRAY)
	values := regex.FindAllStringSubmatch(formatted_evil_array, -1)

	if len(values) > 0 {
		for _, line := range strings.Split(values[0][1], ",") {
			result := tools.Starts_with(line, []string{" "})

			if ok := result[" "]; ok { // It begins with a space
				line = tools.Erase_delimiter(line, []string{" "}, 1)
			}

			line = tools.Erase_delimiter(line, []string{"\""}, -1)
			object.gut = append(object.gut, line)
			object.length++
		}
	}
}

//
//
// Removes the value in the front and returns it
//
//
func (object *Evil_array_t) Pop_front() string {
	to_return := object.gut[0]

	// Move everything one step back
	for i := 1; i < len(object.gut); i++ {
		object.gut[i-1] = object.gut[i]
	}

	object.gut[len(object.gut)-1] = "" // Remove the old element
	object.length--

	return to_return
}

//
//
// Removes the value in the back and returns it
//
//
func (object *Evil_array_t) Pop_back() string {
	to_return := object.gut[len(object.gut)-1]

	object.gut[len(object.gut)-1] = "" // Remove the old element

	object.length--

	return to_return
}

//
//
// Removes the value at the index and returns it
//
//
func (object *Evil_array_t) Pop_index(index int) string {
	to_return := object.Get(index)
	object.gut[index] = "" // Remove the old element

	// Move everything one step back
	for i := index; i < len(object.gut); i++ {
		object.gut[i-1] = object.gut[i]
	}

	object.length--

	return to_return
}

//
//
// Makes the values in the array uppercase
//
//
func (object *Evil_array_t) Uppercase() {
	for i, value := range object.gut {
		object.Replace(i, strings.ToUpper(value))
	}
}

//
//
// Makes the values in the array lowercase
//
//
func (object *Evil_array_t) Lowercase() {
	for i, value := range object.gut {
		object.Replace(i, strings.ToLower(value))
	}
}

package structure

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type Evil_array_t struct {
	gut    []string // The contents of the array
	length int      // The length of the array
}

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
func (object *Evil_array_t) To_string() string {
	toReturn := "${"
	for _, cont := range object.gut {
		toReturn += fmt.Sprintf("\"%s\",", cont)
	}
	toReturn += "}$"
	return toReturn
}

//
//
// Appends data to the array
//
//
func (object *Evil_array_t) Append(new_content string) {
	object.gut = append(object.gut, new_content)
}

//
//
// Replaces data at the index with the new data
//
//
func (object *Evil_array_t) Replace(index int, new_content string) {
	if index < 0 || index > object.Length() {
		notify.Error(fmt.Sprintf("Index %d was out-of-bound", index), "evil_arrays.Replace()")
	}
	object.gut[index] = new_content
}

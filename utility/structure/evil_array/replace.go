package evilarray

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Replaces data at the index with the new data
func (object *Evil_array_t) Replace(index int, new_content string) {
	if index < 0 || index > object.Length() {
		notify.Error(fmt.Sprintf("Index %d was out-of-bound", index), "evil_array.Replace()")
	}
	object.gut[index] = new_content
}

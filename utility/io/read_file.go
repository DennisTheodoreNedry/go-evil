package io

import (
	"io/ioutil"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Reads the contents of the file found in the json data structure
func Read_file(s_json string) string {
	data_object := structure.Receive(s_json)

	content, err := ioutil.ReadFile(data_object.File_path)

	if err != nil {
		notify.Error(err.Error(), "io.Read_file()")
	}

	data_object.File_gut = string(content)

	return structure.Send(data_object)
}

package io

import (
	"io/ioutil"

	"github.com/TeamPhoneix/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Reads the contents of the file found in the json data structure
func Read_file(data_object *json.Json_t) {

	content, err := ioutil.ReadFile(data_object.File_path)

	if err != nil {
		notify.Error(err.Error(), "io.Read_file()")
	}

	data_object.File_gut = string(content)

}

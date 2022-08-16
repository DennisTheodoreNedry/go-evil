package io

import (
	"bufio"
	"os"

	"github.com/TeamPhoneix/go-evil/utility/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

var ALLOWED_EXTENSIONS = [...]string{".evil", ".ge"}

//
//
// Reads the contents of the file found in the json data structure
//
//
func Read_file(s_json string) string {
	data_object := json.Receive(s_json)

	file, err := os.Open(data_object.File_path)

	if err != nil {
		notify.Error(err.Error(), "io.Read_file()")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data_object.Add_file_gut(scanner.Text())
	}

	return json.Send(data_object)
}

//
//
// Compiles the go file into an executable
//
//
func Compile_file(s_json string) string {
	data_object := json.Receive(s_json)
	return json.Send(data_object)

}

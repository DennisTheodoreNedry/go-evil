package io

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Reads the contents of the file found in the json data structure
//
//
func Read_file(s_json string) string {
	data_object := structure.Receive(s_json)

	content, err := ioutil.ReadFile(data_object.File_path)

	if err != nil {
		notify.Error(err.Error(), "io.Read_file()")
	}

	data_object.File_gut = string(content)

	return structure.Send(data_object)
}

//
//
// Writes the malware go content to a local file indicated by the structure
//
//
func Write_file(s_json string) {
	data_object := structure.Receive(s_json)

	_, err := os.Stat(data_object.Malware_path)
	if err != nil {
		if os.IsNotExist(err) {

			if err := os.Mkdir(data_object.Malware_path, 0777); err != nil {
				notify.Error(fmt.Sprintf("Failed to create directory '%s', '%s'", data_object.Malware_path, err.Error()), "io.Write_file")
			}

		} else {
			notify.Error(fmt.Sprintf("Unknown error, %s", err.Error()), "io.Write_file()")
		}
	}

	file, err := os.Create(fmt.Sprintf("%s%s", data_object.Malware_path, data_object.Malware_src_file))

	if err != nil {
		notify.Error(fmt.Sprintf("Failed to open file '%s', '%s'", data_object.Malware_src_file, err.Error()), "io.Write_file()")
	}

	defer file.Close()
	file_stream := bufio.NewWriter(file)

	for _, line := range data_object.Malware_gut {
		if _, err := file_stream.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
			notify.Error(fmt.Sprintf("Failed to write file, %s", err.Error()), "io.Write_file()")
		}
	}

	file_stream.Flush()
}

//
//
// Compiles the go file into an executable
//
//
func Compile_file(s_json string) {
	data_object := structure.Receive(s_json)

	malware := fmt.Sprintf("%s%s%s", data_object.Malware_path, data_object.Binary_name, data_object.Extension)
	src := fmt.Sprintf("%s%s", data_object.Malware_path, data_object.Malware_src_file)

	cmd := exec.Command("go", "build", "-o", malware, "-ldflags=-s -w", src)
	notify.Log("Compiling malware", data_object.Verbose_lvl, "1")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run() // Starts the build

	if err != nil {
		notify.Error(fmt.Sprintf("Failed to compile file, %s", stderr.String()), "io.Compile_file()")
	}

}

//
//
// Compresses the malware
//
//
func Compress_malware(s_json string) {
	data_object := structure.Receive(s_json)
	malware := fmt.Sprintf("%s%s%s", data_object.Malware_path, data_object.Binary_name, data_object.Extension)

	cmd := exec.Command("upx", malware)
	notify.Log("Compressing the malware", data_object.Verbose_lvl, "2")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		notify.Error(fmt.Sprintf("Failed to compress the malware, %s", stderr.String()), "io.Compress_malware()")
	}
}

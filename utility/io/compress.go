package io

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Compresses the malware
func Compress_malware(data_object *json.Json_t) {
	malware := fmt.Sprintf("%s%s%s", data_object.Malware_path, data_object.Binary_name, data_object.Extension)

	cmd := exec.Command("upx", malware)
	data_object.Log_object.Log("Compressing the malware", 2)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		notify.Error(fmt.Sprintf("Failed to compress the malware, %s", err), "io.Compress_malware()", 1)
	}
}

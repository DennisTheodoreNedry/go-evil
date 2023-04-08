package io

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Compresses the malware
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
		notify.Error(fmt.Sprintf("Failed to compress the malware, %s", err), "io.Compress_malware()")
	}
}

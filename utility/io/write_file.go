package io

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Writes the malware go content to a local file indicated by the structure
func Write_file(data_object *json.Json_t) {

	file, err := os.Create(fmt.Sprintf("%s%s", data_object.Malware_path, data_object.Malware_src_file))

	if err != nil {
		notify.Error(fmt.Sprintf("Failed to open file '%s', '%s'", data_object.Malware_src_file, err.Error()), "io.Write_file()", 1)
	}

	defer file.Close()
	file_stream := bufio.NewWriter(file)

	for _, line := range data_object.Malware_gut {
		if _, err := file_stream.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
			notify.Error(fmt.Sprintf("Failed to write file, %s", err.Error()), "io.Write_file()", 1)
		}
	}

	file_stream.Flush()
}

package io

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	mal "github.com/s9rA16Bf4/Malware_Language/domains/malware"
	"github.com/s9rA16Bf4/Malware_Language/utility/notify"
)

func Write_file() {
	example := []string{"package main", "import (sys \"github.com/s9rA16Bf4/Malware_Language/domains/system\"\nwin\"github.com/s9rA16Bf4/Malware_Language/domains/window\"\n)", "func main(){"} // This is always the start of a go program

	example = append(example, mal.Malware_getContent()...) // Insert what the user requested
	example = append(example, "}")                         // And insert the end

	if mal.Malware_getName() == "" {
		mal.Malware_setBinaryName("me_no_virus")
	}

	file, _ := os.Create("output/temp.go") // We utilize a temp directory
	write := bufio.NewWriter(file)

	for _, line := range example {
		_, err := write.WriteString(line + "\n")
		if err != nil {
			notify.Notify_error("Failed to write to disk", "io.write_file()")
		}
	}
	write.Flush()
}

func Read_file(file string) string {
	file_gut, err := ioutil.ReadFile(file)
	if err != nil {
		notify.Notify_error(err.Error(), "io.read_file()")
	}
	return string(file_gut)
}

func Compile_file() {
	arg := "build -o output/" + mal.Malware_getName() + mal.Malware_getExtension() + " output/temp.go"

	fmt.Println(arg)
	cmd := exec.Command("go", strings.Split(arg, " ")...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		notify.Notify_error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
	}
}

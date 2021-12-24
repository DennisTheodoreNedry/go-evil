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

var debug bool = false
var domains []string

func Set_debug(new_debug bool) {
	debug = new_debug
}

const (
	sys  = "\tsys \"github.com/s9rA16Bf4/Malware_Language/domains/system\""
	win  = "\twin \"github.com/s9rA16Bf4/Malware_Language/domains/window\""
	time = "\ttime \"github.com/s9rA16Bf4/Malware_Language/domains/time\""
)

func Append_domain(domain string) {
	if domain == "system" && !find(sys) {
		domains = append(domains, sys)
	} else if domain == "window" && !find(win) {
		domains = append(domains, win)
	} else if domain == "time" && !find(time) {
		domains = append(domains, time)
	}
}

func find(domain string) bool {
	for _, line := range domains {
		if domain == line {
			return true
		}
	}
	return false
}

func Write_file() {
	base_code := []string{
		"package main",
		"import (",
	}
	base_code = append(base_code, domains...)                  // Which domains to include
	base_code = append(base_code, ")", "func main(){")         // Main function and closing include tag
	base_code = append(base_code, "for {")                     // While loop
	base_code = append(base_code, mal.Malware_getContent()...) // Insert the malware code
	base_code = append(base_code, "}}")                        // And insert the end

	if mal.Malware_getName() == "" {
		mal.Malware_setBinaryName("me_no_virus")
	}

	file, _ := os.Create("output/temp.go") // We utilize a temp directory
	write := bufio.NewWriter(file)

	for _, line := range base_code {
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
	cmd := exec.Command("go", strings.Split(arg, " ")...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		notify.Notify_error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
	}

	if !debug {
		arg = "output/temp.go"
		cmd = exec.Command("rm", strings.Split(arg, " ")...)
		err := cmd.Run()

		if err != nil {
			notify.Notify_error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
		}
	}
}

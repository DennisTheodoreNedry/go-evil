package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var domains []string // This will include all domains to copy from our working directory to the output directory

func write_file() {
	example := []string{"package main", "func main(){"} // This is always the start of a go program
	example = append(example, malw.content...)          // Insert what the user requested
	example = append(example, "}")                      // And insert the end

	if malw.malware_name == "" {
		malware_setBinaryName("me_no_virus")
	}

	file, _ := os.Create("output/temp.go") // We utilize a temp directory
	write := bufio.NewWriter(file)

	for _, line := range example {
		_, err := write.WriteString(line + "\n")
		if err != nil {
			notify_error("Failed to write to disk", "io.write_file()")
		}
	}
	write.Flush()

	copy_domains()
}

func read_file(file string) string {
	file_gut, err := ioutil.ReadFile(file)
	if err != nil {
		notify_error(err.Error(), "io.read_file()")
	}
	return string(file_gut)
}

func save_domains(domain string) {
	found := find(domains, domain)
	if !found {
		domains = append(domains, domain)
	}
}

func copy_domains() {
	for _, domain := range domains {
		domain_in, err := os.Open(domain) // Open the domain we want to read from
		if err != nil {
			log.Fatal(err)
		}
		domain_out, err := os.Create("output/" + domain) // Create the domain at the destination
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(domain_out, domain_in) // And copy all content
		if err != nil {
			notify_error(err.Error(), "io.copy_domains()")
		}
	}
}

func compile_file() {
	arg := "build -o output/" + malw.malware_name
	for _, domain := range domains {
		arg += " output/" + domain
	}
	arg += " output/temp.go"

	fmt.Println(arg)
	cmd := exec.Command("go", strings.Split(arg, " ")...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		notify_error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
	}
}

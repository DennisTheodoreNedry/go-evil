package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type malware struct {
	malware_name string   // Malware name
	content      []string // The code that the malware does will be contained here
}

var malw malware

func write_file() {

	if len(malw.content) == 0 {
		example := [...]string{"package main", "import \"fmt\"", "func main(){", "fmt.Println(\"Hello, world\")", "}"}

		for _, line := range example {
			malw.content = append(malw.content, line)
		}
	}
	if malw.malware_name == "" {
		malw.malware_name = "Me_Virus.go"
	}

	file, _ := os.Create(malw.malware_name) // We utilize a temp directory
	write := bufio.NewWriter(file)

	for _, line := range malw.content {
		bytesWritten, _ := write.WriteString(line + "\n")

		fmt.Printf("Written %d\n", bytesWritten)
	}
	write.Flush()

}

func read_file(file string) string {
	file_gut, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(file_gut)
}

func compile_file() {
	out, err := exec.Command("go build " + malw.malware_name).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output %s\n", out)
}

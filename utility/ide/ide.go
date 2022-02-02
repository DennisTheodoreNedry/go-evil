package ide

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/parser"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type ide_t struct {
	file_name   string
	binary_name string
	file_gut    []string
}

var c_ide ide_t

func Main_menu() {
	for {
		fmt.Println("## File information")
		fmt.Printf("File name: %s\n", c_ide.file_name)
		fmt.Println("------------------------------------")
		fmt.Println("## Selection")
		fmt.Println("1. (i)nsert")
		fmt.Println("2. (e)dit")
		fmt.Println("3. (o)pen")
		fmt.Println("4. (r)un")
		fmt.Println("5. (s)ave")
		fmt.Println("6. (q)uit")

		fmt.Print(":: ")
		var user_input string
		fmt.Scanln(&user_input)
		switch user_input {
		case "i", "insert":
			insert()
		case "e", "edit":
			edit()
		case "o", "open":
			open()
		case "r", "run":
			run()
		case "s", "save":
			save()
		case "q", "quit":
			return
		default:
			notify.Inform("Unknown command " + user_input)
		}
	}

}

func insert() {
	for {
		if c_ide.file_name == "" {
			fmt.Print("File name: ")
			fmt.Scanln(&c_ide.file_name)
		} else {
			break
		}
	}
	if c_ide.binary_name == "" {
		fmt.Print("Binary name (optional): ")
		fmt.Scanln(&c_ide.binary_name)
	}

	// Lets take some input
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("################################################")
	for input.Scan() {
		line := input.Text()
		if line == "(quit)" || line == "(q)" {
			break
		} else {
			c_ide.file_gut = append(c_ide.file_gut, line)
		}
	}
}
func edit() {
	if len(c_ide.file_gut) > 0 {
		fmt.Println("################################################")
		for i, line := range c_ide.file_gut {
			if line == "" {
				break
			}
			fmt.Printf("%d] %s\n", i, line)
		}
		fmt.Print("Which line will you edit? or enter '(q)' or '(quit)' to go back")
		var input string
		fmt.Scanln(&input)
		if input == "(q)" || input == "(quit)" {
			return
		} else {
			i := converter.String_to_int(input, "ide.edit()")
			var new_line string
			input := bufio.NewScanner(os.Stdin)
			for input.Scan() {
				line := input.Text()
				if line == "(quit)" {
					break
				} else {
					new_line += line
				}
			}
			c_ide.file_gut[i] = new_line
		}

	} else {
		notify.Inform("Before you can edit a file, make sure that it contains something first")
	}
}
func open() {
	var path string
	fmt.Print("Path to file: ")
	fmt.Scanln(&path)
	file, err := os.Open(path)
	if err != nil {
		notify.Error(err.Error(), "ide.open()")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c_ide.file_gut = append(c_ide.file_gut, scanner.Text())
	}
	index := strings.LastIndex(path, "/")
	c_ide.file_name = path[index+1:]
}

func run() {
	io.Create_file("./output/"+c_ide.file_name, c_ide.file_gut)
	parser.Parser("./output/" + c_ide.file_name)
	io.Write_file()
	io.Compile_file()
	response := io.Run_file("./output/" + mal.GetName())
	fmt.Printf("(resp) %s", response)
	fmt.Println("--------------------------------")
}
func save() {
	io.Create_file("./output/"+c_ide.file_name, c_ide.file_gut)
	notify.Inform("Successfully saved file '" + "./output/" + c_ide.file_name + "'")
}
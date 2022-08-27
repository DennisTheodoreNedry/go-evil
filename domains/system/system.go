package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

//
// Exits the malware
//
func Exit(s_json string, return_code string) (string, string) {
	data_object := structure.Receive(s_json)

	data_object.Add_go_function([]string{
		"func Exit(lvl string){",

		"value := tools.String_to_int(lvl)",

		"os.Exit(value)",

		"}"})

	data_object.Add_go_import("\"github.com/TeamPhoneix/go-evil/utility/tools\"")
	data_object.Add_go_import("\"os\"")

	return fmt.Sprintf("Exit(%s)", return_code), structure.Send(data_object)
}

//
// Prints a message to the screen
//
func Out(s_json string, msg string) (string, string) {
	data_object := structure.Receive(s_json)

	data_object.Add_go_function([]string{
		"func Out(msg string){",

		"fmt.Println(msg)",

		"}"})

	data_object.Add_go_import("\"fmt\"")

	return fmt.Sprintf("Out(%s)", msg), structure.Send(data_object)
}

//
// Executes a command on the running OS and prints the result
//
func Exec(s_json string, cmd string) (string, string) {
	data_object := structure.Receive(s_json)

	data_object.Add_go_function([]string{
		"func Exec(cmd string){",
		"split_cmd := strings.Split(cmd, \" \")",
		"cmd = strings.ReplaceAll(split_cmd[0], \"\\\"\", \"\")",
		"args := strings.ReplaceAll(strings.Join(split_cmd[1:], \" \"), \"\\\"\", \"\")",
		"out, err := exec.Command(cmd, args).Output()",
		"if err != nil {",
		"fmt.Println(err.Error())",
		"}else{",
		"fmt.Println(string(out[:]))",
		"}}"})

	data_object.Add_go_import("\"os/exec\"")
	data_object.Add_go_import("\"fmt\"")
	data_object.Add_go_import("\"strings\"")

	return fmt.Sprintf("Exec(%s)", cmd), structure.Send(data_object)
}

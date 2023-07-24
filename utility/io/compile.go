package io

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
)

// Compiles the go file into an executable
func Compile_file(data_object *json.Json_t) {

	malware := fmt.Sprintf("%s%s%s", data_object.Malware_path, data_object.Binary_name, data_object.Extension)
	src := fmt.Sprintf("%s%s", data_object.Malware_path, data_object.Malware_src_file)
	build_args := []string{}
	compiler := ""

	// Grabs the location of the go environment
	env, err := exec.Command("go", "env", "GOPATH").Output()

	if err != nil {
		notify.Error(err.Error(), "io.Compile_file()", 1)
	}

	go_env := strings.TrimRight(string(env), "\n") // Removes any newline

	// Updates the path variable
	updated_path_env := os.ExpandEnv(fmt.Sprintf("${PATH}:%s/bin", go_env)) // Apparently only provides a formatted string

	if err = os.Setenv("PATH", updated_path_env); err != nil { // So this is needed to *actually* update the path

		notify.Error(err.Error(), "io.Compile_file()", 1)
	}

	// Update the GOOS variable
	if err = os.Setenv("GOOS", data_object.Target_os); err != nil {
		notify.Error(err.Error(), "io.Compile_file()", 1)
	}

	// Update the GOARCH variable
	if err = os.Setenv("GOARCH", data_object.Target_arch); err != nil {
		notify.Error(err.Error(), "io.Compile_file()", 1)
	}

	ldflags := "-ldflags=-s -w"

	if data_object.Obfuscate {
		compiler = "garble"
		build_args = append(build_args, "-literals", "-tiny", "-seed=random", "build", ldflags, "-o", malware, src)
		data_object.Log_object.Log("Compiling malware and obfuscating it, this might take a while", 1)
	} else {
		compiler = "go"
		build_args = append(build_args, "build", "-o", malware, ldflags, src)
		data_object.Log_object.Log("Compiling malware", 1)
	}

	cmd := exec.Command(compiler, build_args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run() // Starts the build

	if err != nil {
		notify.Error(fmt.Sprintf("Failed to compile file, %s\n%s", stderr.String(), err.Error()), "io.Compile_file()", 1)
	}

}

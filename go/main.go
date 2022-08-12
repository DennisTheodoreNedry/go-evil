package main

import (
	"fmt"

	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
)

func main() {
	arg.Argument_add("--file", "-f", true, "Path to the file [REQUIRED]")

	parsed := arg.Argument_parse()

	fmt.Println("Всем привет!")
	fmt.Printf("%s", parsed)
}

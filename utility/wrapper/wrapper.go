package wrapper

import (
	"github.com/TeamPhoneix/go-evil/utility/cleanup"
	"github.com/TeamPhoneix/go-evil/utility/io"
	"github.com/TeamPhoneix/go-evil/utility/parsing"
	"github.com/TeamPhoneix/go-evil/utility/structure/json"
)

// A simple wrapper which combines the process from reading the evil file
// to compiling it. Used by the main process when running gevil from the cli
// but also from when you compile the malware through the text editor
func Parse_and_compile(data_object *json.Json_t) {

	// Parse the file
	parsing.Parse(data_object)

	// Write the file
	io.Write_file(data_object)

	// Compile file
	io.Compile_file(data_object)

	// Compresses the malware
	io.Compress_malware(data_object)

	// Cleanup
	cleanup.Start(data_object)
}

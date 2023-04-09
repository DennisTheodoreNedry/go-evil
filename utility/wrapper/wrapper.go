package wrapper

import (
	"github.com/TeamPhoneix/go-evil/utility/cleanup"
	"github.com/TeamPhoneix/go-evil/utility/io"
	"github.com/TeamPhoneix/go-evil/utility/parsing"
	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// A simple wrapper which combines the process from reading the evil file
// to compiling it. Used by the main process when running gevil from the cli
// but also from when you compile the malware through the text editor
func Parse_and_compile(s_json string) {
	object := structure.Receive(io.Read_file(s_json))
	// Parse the file
	object = structure.Receive(parsing.Parse(structure.Send(object)))

	// Write the file
	io.Write_file(structure.Send(object))

	// Compile file
	io.Compile_file(structure.Send(object))

	// Compresses the malware
	io.Compress_malware(structure.Send(object))

	// Cleanup
	cleanup.Start(structure.Send(object))
}

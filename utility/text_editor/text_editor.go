package text_editor

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/io"
	"github.com/TeamPhoneix/go-evil/utility/structure"
	def_val "github.com/TeamPhoneix/go-evil/utility/text_editor/templates"
	"github.com/TeamPhoneix/go-evil/utility/version"
	"github.com/webview/webview"
)

const (
	TITLE = "** Evil editor **"
)

//
//
// Shows the Main window to the user
//
//
func Spawn_window(s_json string) {
	data_object := structure.Receive(s_json)

	window := webview.New(false)
	defer window.Destroy()
	window.SetTitle(TITLE)

	window.SetSize(data_object.Width, data_object.Height, webview.HintNone)

	data_object = structure.Receive(io.Read_file(structure.Send(data_object)))

	// All the html code for the main webiste
	window.SetHtml(fmt.Sprintf(
		def_val.HTML,
		def_val.CSS,
		def_val.JS_FUNCS,
		TITLE,
		version.TEXT_EDITOR,
		data_object.File_gut, data_object.File_path,
		structure.Send(data_object)))

	window.Bind("Save", Save)
	window.Bind("Compile", Compile)
	window.Run()
}

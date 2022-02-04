package path

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Generate_random_path() string {
	toReturn := ""
	start := "NULL"
	roof := 3 // We will jump three depth in if possible

	switch runtime.GOOS {
	case "windows":
		start = "C:\\"
	default: // Covers linux, mac and everything else
		start = "/"
	}

	for roof > 0 {
		dirs, err := os.ReadDir(start)
		if err != nil {
			notify.Error(err.Error(), "path.Generate_random_path()")
		}
		if len(dirs) == 0 {
			break // We can't continue
		} else {
			tries := 10 // We break after 10 attempts
			for tries > 0 {
				rand.Seed(time.Now().UnixNano())
				next_target := rand.Int() % len(dirs) // From 0 <= x < len(dirs)
				if dirs[next_target].IsDir() {
					toReturn += "/" + dirs[next_target].Name()
					start = toReturn
					break
				}
				tries--
			}
		}
		roof--
	}

	return toReturn + "/"
}

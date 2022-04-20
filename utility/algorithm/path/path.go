package path

import (
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

func Generate_random_path(roof int) string {
	toReturn := ""
	start := "NULL"
	ROOF := roof

	if roof < 1 {
		return start
	}

	switch runtime.GOOS {
	case "windows":
		start = "C:\\"
	default: // Covers linux and mac
		start = "/"
	}

	START := start

	for strings.Count(start, "/") != ROOF {
		roof = ROOF
		start = START

		for roof > 0 {
			dirs, err := os.ReadDir(start)
			if err != nil {
				break
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
	}

	return toReturn + "/"
}

package tools

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
	"github.com/thanhpk/randstr"
)

//
//
// Wrapper for the strings.HasPrefix function, but takes in a array contaning strings to look for
// and returns a map in the format of { "<key>":"<true/false>" }
//
//
func Starts_with(target string, selection []string) map[string]bool {
	to_return := make(map[string]bool)

	for _, value := range selection {
		to_return[value] = strings.HasPrefix(target, value)
	}

	return to_return
}

//
//
// Wrapper for the strings.HasSuffix function, but takes in a array contaning strings to look for
// and returns a map in the format of { "<key>":"<true/false>" }
//
//
func Ends_with(target string, selection []string) map[string]bool {
	to_return := make(map[string]bool)

	for _, value := range selection {
		to_return[value] = strings.HasSuffix(target, value)
	}
	return to_return
}

//
//
// Wrapper for the strings.Contains function, but takes in a array contaning strings to look for
// and returns a map in the format of { "<key>":"<true/false>" }
//
//
func Contains(target string, selection []string) map[string]bool {
	to_return := make(map[string]bool)

	for _, value := range selection {
		to_return[value] = strings.Contains(target, value)
	}
	return to_return
}

//
// Tries to convert the provided string to an int, and returns either
// the converted value or -1 if it failed
//
func String_to_int(value string) int {
	i_value, err := strconv.Atoi(value) // Tries to convert

	if err != nil {
		return -1
	}

	return i_value
}

//
//
// Converts the provided int to a string
//
//
func Int_to_string(value int) string {
	return fmt.Sprint(value)
}

//
//
// Takes the string to work in as a pointer, and makes it go from abc to cba
//
//
func Reverse_string(target *string) {
	output := make([]rune, utf8.RuneCountInString(*target))
	roof := len(output)

	for _, character := range *target {
		roof--
		output[roof] = character
	}
	*target = string(output[0:])
}

//
//
// Splits a string into a string array and returns it
//
//
func Split_string(target string) []string {
	to_return := []string{}
	chars := []rune(target)
	for i := 0; i < len(chars); i++ {
		to_return = append(to_return, string(chars[i]))
	}
	return to_return
}

//
//
// Generates a random string based on the length of the input
//
//
func Generate_random_n_string(size int) string {
	var toReturn string
	for {
		toReturn = randstr.String(size)
		if _, err := strconv.Atoi(string(toReturn[0])); err != nil { // It can't start with a number
			break
		}
	}
	return toReturn
}

//
//
// Generates a random string between the lengthes of 1 and 128
//
//
func Generate_random_string() string {
	rand.Seed(time.Now().UnixNano())
	max := 128
	min := 1
	return Generate_random_n_string(rand.Intn(max-min) + min)
}

//
//
// Generates a random number between min and max
//
//
func Generate_random_int_between(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

//
//
// Generates a random number between 1 and 128
//
//
func Generate_random_int() int {
	return Generate_random_int_between(1, 128)
}

//
//
// Generates a random bool
//
//
func Generate_random_bool() bool {
	value := Generate_random_int_between(1, 2)

	return value == 1
}

//
//
// Erases all occurences of the delimiter in the string
//
//
func Erase_delimiter(line string, delimiters []string, count int) string {

	for _, delimiter := range delimiters {
		line = strings.Replace(line, delimiter, "", count)
	}

	return line
}

//
//
// Returns the username of the current user
//
//
func Grab_username() string {
	user, err := user.Current()

	if err != nil {
		notify.Error(err.Error(), "tools.Grab_username()")
	}

	return user.Username
}

//
//
// Returns the path to the executable file
// including the executable file name
//
//
func Grab_executable_path() string {
	path := os.Args[0]
	return path
}

//
//
// Grabs the current working path
//
//
func Grab_CWD() string {
	path, err := os.Getwd()
	if err != nil {
		notify.Error(err.Error(), "tools.Grab_CWD()")
	}
	return path
}

//
//
// Generates id's from a string utilized for when to reconstruct a string
//
//
func Generate_int_from_string(message string) []int {
	alphabet := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "!", "#", "$", "%", "&", "\"", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", "{", "|", "}", "~", " ", "\t", "\n", "\r", "\x0b", "\x0c"}
	to_return := []int{}

	message = Erase_delimiter(message, []string{"\""}, -1)

	for _, c_msg := range message {
		for id, c_alpha := range alphabet {
			if string(c_msg) == string(c_alpha) {
				to_return = append(to_return, id)
			}
		}
	}

	return to_return
}

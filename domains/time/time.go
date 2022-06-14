package time

import (
	"fmt"
	"time"

	"github.com/s9rA16Bf4/go-evil/utility/converter"
	run_time "github.com/s9rA16Bf4/go-evil/utility/variables/runtime"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type time_t struct { // This is utilized to set when the for loop in `run` will end
	year  int
	month int
	day   int
	hour  int
	min   int
}

var c_time time_t

func SetYear(year string) {
	year = run_time.Check_if_variable(year)
	value := converter.String_to_int(year, "time.SetYear()")
	if value == -1 {
		return
	}
	c_time.year = value
}
func SetMonth(month string) {
	month = run_time.Check_if_variable(month)
	value := converter.String_to_int(month, "time.SetMonth()")
	if value == -1 {
		return
	}
	c_time.month = value
}
func SetDay(day string) {
	day = run_time.Check_if_variable(day)
	value := converter.String_to_int(day, "time.SetDay()")
	if value == -1 {
		return
	}
	c_time.day = value
}
func SetHour(hour string) {
	hour = run_time.Check_if_variable(hour)
	value := converter.String_to_int(hour, "time.SetHour()")
	if value == -1 {
		return
	}
	c_time.hour = value
}
func SetMin(min string) {
	min = run_time.Check_if_variable(min)
	value := converter.String_to_int(min, "time.SetMin()")
	if value == -1 {
		return
	}
	c_time.min = value
}

func Until(value string) {
	if len(value) < 5 || len(value) > 5 {
		notify.Error("Expected format hh:mm", "time.Until()")
	} else {
		SetHour(value[0:2]) // Extracts hour
		SetMin(value[3:5])  // Extracts minute
		Run()
	}
}

func preface() time.Time {
	c_now := time.Now()
	if c_time.year == 0 {
		SetYear(fmt.Sprint(c_now.Year()))
	}
	if c_time.month == 0 {
		SetMonth(fmt.Sprint(int(c_now.Month())))
	}
	if c_time.day == 0 {
		SetDay(fmt.Sprint(c_now.Day()))
	}
	if c_time.hour == 0 {
		SetHour(fmt.Sprint(c_now.Hour()))
	}
	if c_time.min == 0 {
		SetMin(fmt.Sprint(c_now.Minute() + 5)) // Adds five minute just because of the heck of it
	}
	return c_now
}

func Run() {
	c_now := preface()

	for c_time.year != c_now.Year() ||
		c_time.month != int(c_now.Month()) ||
		c_time.day != c_now.Day() ||
		c_time.hour != c_now.Hour() ||
		c_time.min != c_now.Minute() {

		time.Sleep(5 * (10 ^ 9)) // Same as waiting 5 seconds
		c_now = time.Now()       // Update time
		// Could execute some functional value here
	}
}

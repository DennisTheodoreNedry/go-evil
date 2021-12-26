package time

import (
	"time"

	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

type time_t struct { // This is utilized to set when the for loop in `run` will end
	year  int
	month int
	day   int
	hour  int
	min   int
}

var c_time time_t

func Time_setYear(year string) {
	value := converter.String_to_int(year, "time.Time_setYear()")
	c_time.year = value
}
func Time_setMonth(month string) {
	value := converter.String_to_int(month, "time.Time_setMonth()")
	c_time.month = value
}
func Time_setDay(day string) {
	value := converter.String_to_int(day, "time.Time_setDay()")
	c_time.day = value
}
func Time_setHour(hour string) {
	value := converter.String_to_int(hour, "time.Time_setHour()")
	c_time.hour = value
}
func Time_setMin(min string) {
	value := converter.String_to_int(min, "time.Time_setMin()")
	c_time.min = value
}

func Time_until(value string) {
	if len(value) < 5 || len(value) > 5 {
		notify.Notify_error("Expected format hh:mm", "time.Time_until()")
	}

	converter.String_to_int(value[0:2], "time.Time_until()") // Is the hour legit?
	converter.String_to_int(value[3:5], "time.Time_until()") // Is the minutes legit?

	Time_setHour(value[0:2]) // Extracts hour
	Time_setMin(value[3:5])  // Extracts minute
	Time_run()
}

func preface() time.Time {
	c_now := time.Now()
	if c_time.year == 0 {
		c_time.year = c_now.Year()
	}
	if c_time.month == 0 {
		c_time.month = int(c_now.Month())
	}
	if c_time.day == 0 {
		c_time.day = c_now.Day()
	}
	if c_time.hour == 0 {
		c_time.hour = c_now.Hour()
	}
	if c_time.min == 0 {
		c_time.min = c_now.Minute() + 5 // Adds five minute just because of the heck of it
	}
	return c_now
}

func Time_run() {
	c_now := preface()

	for c_time.year != c_now.Year() ||
		c_time.month != int(c_now.Month()) ||
		c_time.day != c_now.Day() ||
		c_time.hour != c_now.Hour() ||
		c_time.min != c_now.Minute() {
		time.Sleep(5 * (10 ^ 9)) // Same as waiting 5 seconds
		c_now = time.Now()       // Update time
	}
}

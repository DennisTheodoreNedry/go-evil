package time

import (
	"fmt"
	"time"

	"github.com/s9rA16Bf4/go-evil/utility/converter"
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

func Time_setYear(year string) {
	value := converter.String_to_int(year, "time.Time_setYear()")
	notify.Log("Setting target year to "+year, notify.Verbose_lvl, "3")
	c_time.year = value
}
func Time_setMonth(month string) {
	value := converter.String_to_int(month, "time.Time_setMonth()")
	notify.Log("Setting target month to "+month, notify.Verbose_lvl, "3")
	c_time.month = value
}
func Time_setDay(day string) {
	value := converter.String_to_int(day, "time.Time_setDay()")
	notify.Log("Setting target day to "+day, notify.Verbose_lvl, "3")
	c_time.day = value
}
func Time_setHour(hour string) {
	value := converter.String_to_int(hour, "time.Time_setHour()")
	notify.Log("Setting target hour to "+hour, notify.Verbose_lvl, "3")
	c_time.hour = value
}
func Time_setMin(min string) {
	value := converter.String_to_int(min, "time.Time_setMin()")
	notify.Log("Setting target min to "+min, notify.Verbose_lvl, "3")
	c_time.min = value
}

func Time_until(value string) {
	if len(value) < 5 || len(value) > 5 {
		notify.Error("Expected format hh:mm", "time.Time_until()")
	}

	Time_setHour(value[0:2]) // Extracts hour
	Time_setMin(value[3:5])  // Extracts minute
	notify.Log("Setting time until execution continues too "+value, notify.Verbose_lvl, "2")
	Time_run()
}

func preface() time.Time {
	c_now := time.Now()
	if c_time.year == 0 {
		Time_setYear(fmt.Sprint(c_now.Year()))
	}
	if c_time.month == 0 {
		Time_setMonth(fmt.Sprint(c_now.Month()))
	}
	if c_time.day == 0 {
		Time_setDay(fmt.Sprint(c_now.Day()))
	}
	if c_time.hour == 0 {
		Time_setHour(fmt.Sprint(c_now.Hour()))
	}
	if c_time.min == 0 {
		Time_setMin(fmt.Sprint(c_now.Minute() + 5)) // Adds five minute just because of the heck of it
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
		// Could execute some functional value here
	}
}

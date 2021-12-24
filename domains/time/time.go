package time

import (
	"strconv"
	"time"

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
	value, err := strconv.Atoi(year)
	if err != nil {
		notify.Notify_error("Failed to convert "+year+" into integer", "time.until()")
	}
	c_time.year = value
}
func Time_setMonth(month string) {
	value, err := strconv.Atoi(month)
	if err != nil {
		notify.Notify_error("Failed to convert "+month+" into integer", "time.until()")
	}
	c_time.month = value
}
func Time_setDay(day string) {
	value, err := strconv.Atoi(day)
	if err != nil {
		notify.Notify_error("Failed to convert "+day+" into integer", "time.until()")
	}
	c_time.day = value
}
func Time_setHour(hour string) {
	value, err := strconv.Atoi(hour)
	if err != nil {
		notify.Notify_error("Failed to convert "+hour+" into integer", "time.until()")
	}
	c_time.hour = value
}
func Time_setMin(min string) {
	value, err := strconv.Atoi(min)
	if err != nil {
		notify.Notify_error("Failed to convert "+min+" into integer", "time.until()")
	}
	c_time.min = value
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

	}
}

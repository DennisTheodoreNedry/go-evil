package time

import (
	"time"
)

type time_t struct { // This is utilized to set when the for loop in `run` will end
	year  int
	month int
	day   int
	hour  int
	min   int
}

var c_time time_t

func Time_setYear(year int) {
	c_time.year = year
}
func Time_setMonth(month int) {
	c_time.month = month
}
func Time_setDay(day int) {
	c_time.day = day
}
func Time_setHour(hour int) {
	c_time.hour = hour
}
func Time_setMin(min int) {
	c_time.min = min
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
		c_time.min = 00
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

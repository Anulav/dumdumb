package utils

import "time"

func SleepMinutes(min int) {
	time.Sleep(time.Duration(min))
}

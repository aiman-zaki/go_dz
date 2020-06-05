package wrappers

import (
	"fmt"
	"time"
)

//StringToTime : Convert string to time.Time
func StringToTime(timeString string) time.Time {
	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

package main

import (
	"fmt"
	"strconv"
	"time"
)

func convertElapsedSecondsToString(seconds int64) string {
	time_elapsed_str := strconv.FormatInt(seconds, 10) + "s"
	d, err := time.ParseDuration(time_elapsed_str)
	if err != nil {
		panic(err)
	}
	hour := int64(d/time.Hour)
	min := int64(d/time.Minute) - hour * 60
	sec := seconds - hour * 60 * 60 - min * 60
	h := strconv.FormatInt(hour, 10)
	m := strconv.FormatInt(min, 10)
	s := strconv.FormatInt(sec, 10)
	return fmt.Sprintf("%s:%s:%s", h, m, s)
}

func main() {
	timeElapsed := convertElapsedSecondsToString(20555)
	fmt.Println(timeElapsed)
}
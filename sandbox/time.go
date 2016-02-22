package main

import (
    "fmt"
    "strconv"
    "time"
)

func convertElapsedSecondsToString(seconds int64) string {
    timeElapsedString := strconv.FormatInt(seconds, 10) + "s"
    duration, err := time.ParseDuration(timeElapsedString)
    if err != nil {
        panic(err)
    }
    hour := int64(duration/time.Hour)
    min := int64(duration/time.Minute) - hour * 60
    sec := seconds - hour * 60 * 60 - min * 60
    h := strconv.FormatInt(hour, 10)
    m := strconv.FormatInt(min, 10)
    s := strconv.FormatInt(sec, 10)
    return fmt.Sprintf("%s:%s:%s", h, m, s)
}

func main() {
    myTestRunDuration := convertElapsedSecondsToString(20555)
    fmt.Println(myTestRunDuration)
}
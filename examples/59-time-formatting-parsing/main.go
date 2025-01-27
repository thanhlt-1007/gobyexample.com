package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Printf("now: %v\n", now)
    fmt.Printf("Format RFC1123: %s\n", now.Format(time.RFC1123))
    fmt.Printf("Format RFC3339: %s\n", now.Format(time.RFC3339))

    parsedTime, _ := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00",
    )
    fmt.Printf("parsedTime: %v\n", parsedTime)

    fmt.Printf("Format 3:04PM: %v\n", now.Format("3:04PM"))
    fmt.Printf("Format Mon Jan _2 15:04:05 2006: %v\n", now.Format("Mon Jan _2 15:04:05 2006"))
    fmt.Printf("Format 2006-01-02T15:04:05.999999-07:00: %v\n", now.Format("2006-01-02T15:04:05.999999-07:00"))

    format := "3 04 PM"
    parsedTime, _ = time.Parse(format, "8 41 PM")
    fmt.Printf("parsedTime: %v\n", parsedTime)

    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        now.Year(), now.Month(), now.Day(),
        now.Hour(), now.Minute(), now.Second())

    ansic := "Mon Jan _2 15:04:05 2006"
    _, err := time.Parse(ansic, "8:41PM")
    fmt.Printf("err: %v\n", err)
}

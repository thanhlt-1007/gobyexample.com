package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Printf("now: %v\n", now)

    date := time.Date(
        2025,
        12,
        31,
        12,
        30,
        00,
        00,
        time.UTC,
    )
    fmt.Printf("date: %v", date)
    fmt.Printf("date.Year: %d\n", date.Year())
    fmt.Printf("date.Month: %d\n", date.Month())
    fmt.Printf("date.Day: %d\n", date.Day())
    fmt.Printf("date.Hour: %d\n", date.Hour())
    fmt.Printf("date.Minute: %d\n", date.Minute())
    fmt.Printf("date.Second: %d\n", date.Second())
    fmt.Printf("date.Nanosecond: %d\n", date.Nanosecond())
    fmt.Printf("date.Location: %v\n", date.Location())
    fmt.Printf("date.Weekday: %v\n", date.Weekday())

    fmt.Printf("date.Before: %t\n", date.Before(now))
    fmt.Printf("date.After: %t\n", date.After(now))
    fmt.Printf("date.Equal: %t\n", date.Equal(now))

    diff := now.Sub(date)
    fmt.Printf("diff: %v\n", diff)
    fmt.Printf("diff.Hours: %f\n", diff.Hours())
    fmt.Printf("diff.Minutes: %f\n", diff.Minutes())
    fmt.Printf("diff.Seconds: %f\n", diff.Seconds())
    fmt.Printf("diff.Nanoseconds: %d\n", diff.Nanoseconds())

    fmt.Printf("date.Add %v\n", date.Add(diff))
}

package main

import (
    "fmt"
    "regexp"
)

func main() {
    matched, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Printf("matched: %t\n", matched)

    re, _ := regexp.Compile("p([a-z]+)ch")
    fmt.Printf("matched: %t\n", re.MatchString("peach"))

    fmt.Printf("FindString: %s\n", re.FindString("peach punch"))
    fmt.Println("FindStringIndex:", re.FindStringIndex("peach punch"))
    fmt.Println("FindStringSubmatch:", re.FindStringSubmatch("peach punch"))
    fmt.Println("FindStringSubmatchIndex:", re.FindStringSubmatchIndex("peach punch"))
}

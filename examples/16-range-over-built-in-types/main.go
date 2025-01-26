package main

import "fmt"

func main() {
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for index, num := range nums {
        if num == 3 {
            fmt.Println("index:", index)
        }
    }

    fruits := map[string]string {
        "a": "apple",
        "b": "banana",
    }
    for k, fruit := range fruits {
        fmt.Printf("%s -> %s\n", k, fruit)
    }
    for k := range fruits {
        fmt.Printf("%s\n", k)
    }

    for index, char := range "go" {
        fmt.Printf("%d -> %c\n", index, char)
    }
}

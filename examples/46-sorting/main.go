package main

import (
    "fmt"
    "slices"
)

func main() {
    strings := []string {"c", "a", "b"}
    fmt.Println("strings:", strings)
    slices.Sort(strings)
    fmt.Println("strings:", strings)

    ints := []int {7, 2, 4}
    fmt.Println("ints:", ints)
    slices.Sort(ints)
    fmt.Println("ints:", ints)

    fmt.Println("IsSorted:", slices.IsSorted(ints))
}

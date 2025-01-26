package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    s := 0
    for _, num := range nums {
        s += num
    }
    fmt.Println(s)
}

func main() {
    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}

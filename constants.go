package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println("s:", s)

    const n = 500000000
    fmt.Println("n:", n)

    const d = 3e20 / n
    fmt.Println("d:", int64(d))

    fmt.Println(math.Sin(n))
}

package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

	nextInt2 := intSeq()
    fmt.Println(nextInt2())
}

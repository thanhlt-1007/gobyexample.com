package main

import "fmt"

func zeroValue(i int) {
    i = 0
}

func zeroPointer(i *int) {
    *i = 0
}

func main() {
    i := 1
    fmt.Printf("initial: %d\n", i)

    zeroValue(i)
    fmt.Printf("zeroValue: %d\n", i)

    zeroPointer(&i)
    fmt.Printf("zeroPointer: %d\n", i)
}

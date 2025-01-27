package main

import "fmt"

func raisePanic() {
    panic("a problem")
}

func recoverPanic() {
    err := recover()
    if err != nil {
        fmt.Printf("Recoverd. Error: %v\n", err)
    }
}

func main() {
    defer recoverPanic()

    raisePanic()

    // not executed
    fmt.Println("After raisePanic()")
}

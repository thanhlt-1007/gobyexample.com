package main

import (
    "fmt"
    "os"
)

type Point struct {
    X int
    Y int
}

func main() {
    point := Point{
        X: 1,
        Y: 2,
    }
    fmt.Printf("point: %v\n", point)
    fmt.Printf("point: %+v\n", point)
    fmt.Printf("point: %#v\n", point)
    fmt.Printf("point: %T\n", point)

    fmt.Printf("bool: %t\n", true)
    fmt.Printf("int: %d\n", 123)
    fmt.Printf("bin: %b\n", 14)
    fmt.Printf("char: %c\n", 33)
    fmt.Printf("hex: %x\n", 456)

    fmt.Printf("float: %f\n", 78.9)
    fmt.Printf("float: %e\n", 78.9)
    fmt.Printf("float: %E\n", 78.9)

    fmt.Printf("str: %s\n", "\"string\"")
    fmt.Printf("str: %q\n", "\"string\"")
    fmt.Printf("str: %x\n", "hex this")

    fmt.Printf("pointer: %p\n", &point)

    fmt.Printf("width: |%6d|%6d|\n", 12, 345)
    fmt.Printf("width: |%6.2f|%6.2f|\n", 1.2, 3.45)
    fmt.Printf("width: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
    fmt.Printf("width: |%6s|%6s|\n", "foo", "bar")
    fmt.Printf("width: |%-6s|%-6s|\n", "foo", "bar")

    str := fmt.Sprintf("Sprintf: a %s", "string")
    fmt.Println(str)

    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}

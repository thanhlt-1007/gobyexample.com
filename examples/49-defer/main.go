package main

import (
    "fmt"
    "os"
)

func createFile(path string) *os.File {
    fmt.Println("creating")
    file, err := os.Create(path)
    if err != nil {
        panic(err)
    }

    return file
}

func writeFile(file *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
    fmt.Println("closing")
    err := file.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func main() {
    file := createFile("/tmp/defer.txt")
    writeFile(file)
    defer closeFile(file)
}

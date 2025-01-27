package main

import (
    "errors"
    "fmt"
)

type ArgError struct {
    Arg int
    Message string
}

func (ergError ArgError) Error() string {
    return fmt.Sprintf("%d - %s", ergError.Arg, ergError.Message)
}

func error42(num int) (int, error) {
    if num == 42 {
        return -1, ArgError{
            Arg: num,
            Message: "num is 42",
        }
    }

    return num, nil
}

func main() {
    _, err := error42(42)
    var argError ArgError
    if errors.As(err, &argError) {
        fmt.Println("Error match ArgError")
        fmt.Printf("Arg: %d\n", argError.Arg)
        fmt.Printf("Message: %s\n", argError.Message)
    } else {
        fmt.Println("Error doesn't match ArgError")
    }
}

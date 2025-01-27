package main

import (
    "errors"
    "fmt"
)

func error42(num int) (int, error) {
    if num == 42 {
        return -1, errors.New("num is 42")
    }

    return num, nil
}

func makeTea(tea int) error {
    if tea == 2 {
        return ErrorOutOfTea
    } else if tea == 3 {
        return fmt.Errorf("making tea: %w", ErrorOutOfPower)
    }

    return nil
}

var ErrorOutOfTea = fmt.Errorf("out of tea")
var ErrorOutOfPower = fmt.Errorf("out of power")

func main() {
    nums := []int {7, 42}
    for _, num := range nums{
        _, err := error42(num)
        if err != nil {
            fmt.Println("Failed with num", num, "error", err)
        } else {
            fmt.Printf("Word with num %d\n", num)
        }
    }

    for i := range 10 {
        err := makeTea(i)
        if err != nil {
            if errors.Is(err, ErrorOutOfTea) {
                fmt.Printf("i: %d: We should buy new tea!\n", i)
            } else if errors.Is(err, ErrorOutOfPower) {
                fmt.Printf("i: %d: Now it is dark!\n", i)
            }

            continue
        }
        fmt.Printf("i: %d Tea is ready!\n", i)
    }
}

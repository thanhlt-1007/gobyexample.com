package main

import (
    "fmt"
    "strconv"
)

func main() {
    strToF, _ := strconv.ParseFloat("1.234", 64)
    fmt.Printf("Float: %f\n", strToF)

    strToI, _ := strconv.ParseInt("123", 0, 64)
    fmt.Printf("Int: %d\n", strToI)

    hexToI, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Printf("Int: %d\n", hexToI)

    strToU, _ := strconv.ParseUint("789", 0, 64)
    fmt.Printf("Unit: %d\n", strToU)

    aToI, _ := strconv.Atoi("135")
    fmt.Printf("Int: %d\n", aToI)

    _, err := strconv.Atoi("wat")
    fmt.Printf("Error: %v\n", err)
}

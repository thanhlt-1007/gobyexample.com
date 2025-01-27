package main

import (
    "fmt"
    "math/rand/v2"
)

func main() {
    fmt.Printf("%d\n", rand.IntN(100))
    fmt.Printf("%d\n", rand.IntN(100))

    fmt.Printf("%f\n", rand.Float64())
    fmt.Printf("%f\n", rand.Float64())

    pcg := rand.NewPCG(42, 1024)
    pcgRand := rand.New(pcg)
    fmt.Printf("%d\n", pcgRand.IntN(100))
    fmt.Printf("%d\n", pcgRand.IntN(100))
}

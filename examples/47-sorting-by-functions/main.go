package main

import (
    "cmp"
    "fmt"
    "slices"
)

type Person struct {
    Name string
    Age int
}

func CompareLen(a string, b string) int {
    return cmp.Compare(len(a), len(b))
}

func CompareAge(a Person, b Person) int {
    return cmp.Compare(a.Age, b.Age)
}

func main() {
    fruits := []string {
        "peach",
        "banana",
        "kiwi",
    }

    fmt.Println("fruits:", fruits)
    slices.SortFunc(fruits, CompareLen)
    fmt.Println("fruits:", fruits)

    persons := []Person {
        Person {
            Name: "Jax",
            Age: 37,
        },
        Person {
            Name: "TJ",
            Age: 25,
        },
        Person {
            Name: "Alex",
            Age: 72,
        },
    }
    fmt.Println("persons:", persons)
    slices.SortFunc(persons, CompareAge)
    fmt.Println("persons:", persons)
}

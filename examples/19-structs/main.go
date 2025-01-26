package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func newPerson(name string) *Person {
    person := Person {
        Name: name,
        Age: 42,
    }

    return &person
}

func main() {
    bob := Person{"Bob", 20}
    fmt.Println(bob)

    alice := Person {
        Name: "Alice",
        Age: 20,
    }
    fmt.Println(alice)

    fred := Person {
        Name: "Fred",
    }
    fmt.Println(fred)

    ann := Person {
        Name: "Ann",
        Age: 40,
    }
    fmt.Println(&ann)

    john := *newPerson("John")
    fmt.Println(john)

    sean := Person {
        Name: "Sean",
        Age: 50,
    }
    fmt.Println(sean.Name)

    sp := sean
    fmt.Println(sp.Age)

    sean.Age = 51
    fmt.Println(sean.Age)
    fmt.Println(sp.Age)
}

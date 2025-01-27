package main

import "fmt"

type Base struct {
    Num int
}

func (base Base) describe() string {
    return fmt.Sprintf("Base with Num = %d", base.Num)
}

type Container struct {
    Base
    Str string
}

type Describeable interface {
    describe() string
}

func main() {
    base := Base{
        Num: 1,
    }
    container := Container{
        Base: base,
        Str: "CONTAINER",
    }

    fmt.Printf("base.Num: %d\n", base.Num)
    fmt.Printf("container.Base.Num: %d\n", container.Base.Num)
    fmt.Printf("container.Num: %d\n", container.Num)

    fmt.Printf("base.describe(): %s\n", base.describe())
    fmt.Printf("container.base.describe(): %s\n", container.Base.describe())
    fmt.Printf("container.describe(): %s\n", container.describe())
}

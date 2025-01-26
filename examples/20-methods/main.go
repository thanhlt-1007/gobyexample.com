package main

import "fmt"

type Rect struct {
    Width, Height, Area int
}

func (rect Rect) area() int {
    area := rect.Width * rect.Height
    // not update rect
    rect.Area = area
    return area
}

func (rect *Rect) setArea() int {
    area := rect.Width * rect.Height
    // update rect
    rect.Area = area
    return area
}

func main() {
    rect := Rect {
        Width: 20,
        Height: 10,
    }
    fmt.Printf("area(): %d\n", rect.area())
    fmt.Printf("rect.Area: %d\n", rect.Area)

    fmt.Printf("setArea(): %d\n", rect.setArea())
    fmt.Printf("rect.Area: %d\n", rect.Area)
}

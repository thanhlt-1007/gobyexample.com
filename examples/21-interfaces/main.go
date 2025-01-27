package main

import (
    "fmt"
    "math"
)

type Rect struct {
    Width, Height float64
}

func (rect Rect) perim() float64 {
    return 2 * rect.Width + 2 *rect.Height
}

func (rect Rect) area() float64 {
    return rect.Width * rect.Height
}

type Circle struct {
    Radius float64
}

func (circle Circle) perim() float64 {
    return 2 * math.Pi * circle.Radius
}

func (circle Circle) area() float64 {
    return math.Pi * circle.Radius * circle.Radius
}

type Geometry interface {
    area() float64
    perim() float64
}

func measure(geometry Geometry) {
    fmt.Println(geometry)
    fmt.Printf("area: %f\n", geometry.area())
    fmt.Printf("perim: %f\n", geometry.perim())
}

func isCircle(geometry Geometry) bool {
    circle, ok := geometry.(Circle)
    if ok {
        fmt.Printf("Is Circle with radius: %f\n", circle.Radius)
        return true
    } else {
        fmt.Println("Isn't Circle")
        return false
    }
}

func main() {
    rect := Rect {
        Width: 3,
        Height: 4,
    }
    measure(rect)
    isCircle(rect)

    circle := Circle{10}
    measure(circle)
    isCircle(circle)
}

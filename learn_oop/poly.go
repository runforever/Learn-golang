package main

import (
        "fmt"
)

type Square struct {
        width float64
}

type Circle struct {
        redius float64
}

func (s Square) Area() float64 {
        return s.width * s.width
}

func (s Circle) Area() float64 {
        return s.redius * s.redius * 3.14
}

type PolySharp interface {
        Area() float64
}

func main() {
        var s1 PolySharp = &Square{4}
        var s2 PolySharp = &Circle{4}
        fmt.Println(s1.Area())
        fmt.Println(s2.Area())
}

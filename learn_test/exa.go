package main

import "fmt"

func Add(x, y int) int {
        return x + y
}

func main() {
        x, y := 1, 2
        fmt.Println(Add(x, y))
}

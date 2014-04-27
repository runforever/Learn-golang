package main

import "fmt"

func byref(a int) {
        a += 5
        fmt.Println(a)
}

func byval(a *int) {
        *a += 5
        fmt.Println(*a)
}

func main() {
        a := 4
        fmt.Println(a)
        byref(a)
        fmt.Println(a)
        byval(&a)
        fmt.Println(a)
}

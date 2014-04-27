package main

import "fmt"

func Add(ch chan int, x, y int) {
        z := x + y
        // 往channel中写数据
        ch <- z
}

func main() {
        // 定义10个channel
        chs := make([]chan int, 10)

        for i := 0; i < 10; i++ {
                chs[i] = make(chan int)
                // 使用goroutine
                go Add(chs[i], i, 1)
        }

        for i, ch := range chs {
                fmt.Println("chan index", i)
                // 读channel中的数据
                value := <-ch
                fmt.Println("value", value)
        }
}

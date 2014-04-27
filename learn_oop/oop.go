package main

import "fmt"

/* 定义一个形状类
   实现正方形计算面积的方法
   实现圆计算面积的方法
*/
type Sharp struct {
        r float64
}

func (s Sharp) Square() float64 {
        return s.r * s.r
}

func (s Sharp) Circle() float64 {
        return s.r * s.r * 3.14
}

type Area interface {
        Square() float64
        Circle() float64
}

func main() {
        // 实例化类方法, 引用
        var sharpSquare Sharp = Sharp{4}
        var sharpCircle Sharp = Sharp{4}

        // 实例化类方法，指针
        var sharpSquare1 *Sharp = &Sharp{4}
        var sharpCircle1 *Sharp = &Sharp{4}

        // 实例化类方法, 使用new关键字
        // new 分配了0值填充的Sharp类型的内存空间，并且返回该空间的地址
        sharpSquare2 := new(Sharp)
        sharpSquare2.r = 5

        fmt.Println("引用")
        fmt.Println(sharpSquare.Square())
        fmt.Println(sharpCircle.Circle())

        fmt.Println("指针")
        fmt.Println(sharpSquare1.Square())
        fmt.Println(sharpCircle1.Circle())

        fmt.Println("new关键字")
        fmt.Println(sharpSquare2.Circle())

        // Sharp 实现了接口的两个方法，已经实现了该接口
        // 接口调用
        fmt.Println("接口")
        var area1 Area = Sharp{6}
        fmt.Println(area1.Square())
        fmt.Println(area1.Circle())
}

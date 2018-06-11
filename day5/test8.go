// Q5. (0) 平均值
// 1. 编写一个函数用于计算一个 float64 类型的 slice 的平均值。

package main

import "fmt"

func getArea(num [7]float64) float64 {
    var sum float64
    for _ , v := range num {
        sum += v
    }
    
    return sum / float64(len(num))
}


func main() {
    var num  = [...]float64{1,2,3,4,1,3,1}
    var area float64 = getArea(num)
    fmt.Println(area)
}

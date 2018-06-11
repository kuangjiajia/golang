// Q12. (0) 最小值和最大值
// 1. 编写一个函数，找到 int slice ([]int) 中的最大值。
// 2. 编写一个函数，找到 int slice ([]int) 中的最小值。

package main

import "fmt"

func findBiggest(num []int) int {
    var tmp int 
    for _,v := range num {
        if v > tmp {
            tmp = v
        }
    }
    return tmp
}

func findLeast(num []int) int {
    var tmp int = num[0]
    for _,v := range num {
        if v < tmp {
            tmp = v
        }
    }
    return tmp
}

func main() {
    var k = []int{-31,2,3,4,5,7,6}
    x := findBiggest(k)
    fmt.Println(x)
    l := findLeast(k)
    fmt.Println(l)
}
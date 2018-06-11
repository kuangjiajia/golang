// Q6. (0) 整数顺序
// 1. 编写函数，返回其（两个）参数正确的（自然）数字顺序：
// f(7,2) → 2,7
// f(2,7) → 2,7

package main

import "fmt"

func getOrder(num1 int,num2 int) (int,int){
    if num1 <= num2 {
        return num1 , num2
    }
    return num2 , num1
}
 

func main() {
    num1, num2 := getOrder(6,4)
    fmt.Println(num1)
    fmt.Println(num2)
}
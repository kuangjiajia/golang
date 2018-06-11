// 斐波拉切数列

package main

import "fmt"

func fib(num int) int{
    if num == 0 {
        return 0
    }else if num == 1 {
        return 1
    }else {
        return fib(num -1) + fib(num -2)
    }
}

func main() {
    var arr [40]int 
    for i := 0 ; i < 40 ; i++ {
        arr[i] = fib(i)
    }
    fmt.Println(arr)
}
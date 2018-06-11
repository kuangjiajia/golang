// Q3. (1) 字符串
// 1. 建立一个 Go 程序打印下面的内容（到 100 个字符）：
// A
// AA
// AAA
// AAAA
// AAAAA
// AAAAAA
// AAAAAAA

package main

import "fmt"

func main() {
    for i := 1 ; i <= 100 ; i++ {
        for j := 0 ; j < i ; j++ {
            fmt.Print("A")
        }
        fmt.Println("")
    }
}
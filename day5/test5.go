// 3. 扩展/修改上一个问题的程序，替换位置 4 开始的三个字符为 “abc”。

package main

import ("fmt"

       )


func main() {
    s := "asSASA ddd dsjkdsjs dk"
    s = s[:4] + "abc"
    fmt.Println(s)
}
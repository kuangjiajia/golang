// 递归


package main

import ("fmt"
        )

func dig(x int) {
    if x == 10 {
        return
    }
    x++
    dig(x)
    fmt.Println(x)
}
 
func main() {
   dig(0)
}
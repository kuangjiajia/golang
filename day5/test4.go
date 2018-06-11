// 2. 建立一个程序统计字符串里的字符数量：
// asSASA ddd dsjkdsjs dk
// 同时输出这个字符串的字节数。提示： 看看 unicode/utf8 包。

package main

import "fmt"
import "unicode/utf8"


func main() {
    s := "匡asSASA ddd dsjkdsjs dk"
    var sum int 
    for i := 0 ; i < len(s); {
        r, n := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%c：%v   ", r, n) 
		i += n
		sum += n
    }
    fmt.Println(sum)
}
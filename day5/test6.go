// 编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”。提示：不
// 幸的是你需要知道一些关于转换的内容，参阅 “转换” 第 59 页的内容

package main

import ("fmt"
        )

func reverseStr(str string) string {
    runeStr := []rune(str)
    for i,j := 0,len(runeStr)-1 ; j > i ; j,i = j-1,i+1 {
        runeStr[i],runeStr[j] = runeStr[j],runeStr[i]
    }
    return string(runeStr)
}

func main() {
    s := "raboof"
    x := reverseStr(s)
    fmt.Println(x)
}
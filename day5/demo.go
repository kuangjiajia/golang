package main

import "fmt"

func main() {
   s := "kuangajiajia"
   arr := []rune(s)
   arr[0] = 'l'
   fmt.Println(string(arr))
   const name int = 123
   if name > 2 {
       fmt.Println(456)
   }
   var x int = 5
llp:
    if x > 0 {
        fmt.Println(x)
        x--
        goto llp
    }
    users := []string{"kuangjiajia","liliping","zhuzhangxuan"}
    for index,value := range users {
        fmt.Println(index)
        fmt.Println(value)
    }
    const i int = 10
    switch {
        case i > 8: 
            fmt.Println("yes")
            fallthrough
        case i > 9:
            fmt.Println("no")
    }
}
package main

import ("fmt"
        )

var a int 
func f() {
    a = 10
    fmt.Println("f")
    fmt.Println(a)
    g()
}

func g() {
    a := 30
    fmt.Println("g")
    fmt.Println(a)
}

func demo() (n int , str string) {
    x := 20
    n += x 
    str = "hashdsahdhashd"
    return 
}

func test() {
    defer fmt.Println("end")
    fmt.Println("start")
}
 
func rangeInt(args ...int) {
    // for i,v := range args {
    //     fmt.Println(v)
    //     fmt.Println(i)
    // }
    range2Int(args...)
}
 
func printName(name string) {
    fmt.Println(name)
}
func personPrint(name string, f func(n string)) {
    f(name)
}
 
func range2Int(args ...int) {
    for i,v := range args {
        fmt.Println(v)
        fmt.Println(i)
    }
}


func main() {
   f()
   age,name := demo()
   fmt.Println(age)
   fmt.Println(name)
   test()
   func (x int) {
       fmt.Println(x)
   }(10)
   rangeInt(10,11,12,13)
   a := func() {
       fmt.Println("func")
   }
   a() 
   var wx = map[int]func() int {
       1: func() int {
           return 123
       },
   }
   fmt.Println(wx)
   personPrint("kuangjiajia",printName)
}
package main

import (
	"fmt"
)

// import (
// 	"fmt"
// )


func main() {
	// if x := 20; x > 20 {
	// 	fmt.Println("x is bigger than 20")
	// }else {
	// 	fmt.Println("x is smaller than 20")		
	// }
	fmt.Println("x is smaller than 20")		
	forFunc()
}

// func myFunc() {
// 	i := 0
// Here:   //这行的第一个词，以冒号结束作为标签
// 	println(i)
// 	i++
// 	goto Here   //跳转到Here去
// } 

func forFunc() {
	// m := 1
	// fmt.Println(1)
	// for m < 10 {
	// 	m += m
	// }
	// fmt.Println(2)
	for k,v:=range map {
		fmt.Println("map's key:",k)
		fmt.Println("map's val:",v)
	}
}
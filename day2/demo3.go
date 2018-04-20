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
	// fmt.Println("x is smaller than 20")		
	// forFunc()
	// rangeFunc()
	// switchFunc(2)
	x,y := returnFunc(1,3)
	fmt.Print(x,y)
}

// func myFunc() {
// 	i := 0
// Here:   //这行的第一个词，以冒号结束作为标签
// 	println(i)
// 	i++
// 	goto Here   //跳转到Here去
// } 

// func forFunc() {
	// m := 1
	// fmt.Println(1)
	// for m < 10 {
	// 	m += m
	// }
	// fmt.Println(2)

// }



// func switchFunc(val int) {
// 	switch val {
// 	case 2:
// 		fmt.Println("2")
// 	case 3:
// 		fmt.Println("3")
// 	default:
// 		fmt.Println("no")
// 	}
// }

// func rangeFunc() {
// 	kjj := map[int]int{1: 3,2: 5}
// 	for v,i := range kjj {
// 		fmt.Printf("v : %d, i : %d ",v,i)
// 	}
// }



// func returnFunc(i int,j int) int{
// 	return i+j
// }

// func returnFunc(i int,j int)(int,int) {
// 	return i+j,j*i
// }
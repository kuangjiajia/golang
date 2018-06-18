package main

import "fmt"

func addr() func(int) int {
	sum := 0
	return func(i int) int{
		sum += i
		return sum
	}
}

type iAdder func(int) (int,iAdder)
//
//func adder2(base int) iAdder {
//	return func(v int) (int ,iAdder) {
//		return base + v , adder2(base+v)
//	}
//}

//type iAdder func(int) (int,iAdder) {}

func adder2(base int) iAdder  {
	return func(v int) (int,iAdder) {
		return base+v,adder2(base+v)
	}
}


func main() {
	//a := addr() //返回的函数
	a := adder2(0)
	for j := 0 ; j < 10 ; j++ {
		//fmt.Println(a(j))
		//x,_ := kjj(j)
		var s int
		s,a = a(j)
		fmt.Println(s)
	}

}
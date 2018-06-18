package main

import (
	"fmt"
)

type Person struct {
	name string "kuangjiajia"
	age int
}

type Add func(i,j int) int

func add(x,y int) int{
	return x+y
}

func cacl(i,j int,kjj Add) int{
	result := kjj(i,j)
	return result
}



//func show(i interface{}) {
//	switch i.(type) {
//	case *Person:
//		t := reflect.TypeOf(i)
//		v := reflect.ValueOf(i)
//		tag := t.Elem().Field(0).Tag
//		name : v.Elem().Field(0).String()
//	}
//}

//func main() {
//	fmt.Println(cacl(1,3,add))
//	x := "asd"
//	fn := func() {
//		fmt.Println(x)
//	}
//	fn()
//	type funcType func()
//	fn2 := fn
//	fn2()
//}

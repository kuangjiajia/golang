package main


//import (
//	"fmt"
//	"reflect"
//)
//
//type person struct {
//	name string
//	age int
//}
//
//type course []string
//
//type student struct {
//	name string
//	grade string
//	age int
//}
//
//type goods struct {
//	name string "goods_name"
//	price string "goods_price"
//}
//
//type kjj struct {
//	jjk student
//	hobby string
//	course
//}
//
//type funcs struct {
//	radius float64
//}
//
//func (f funcs) getArea() float64 {
//	return f.radius * f.radius
//}
//
//
//
//
//
//func main() {
//	var p1 person
//	p1.name = "kuangjiajia"
//	p1.age = 12
//	fmt.Println(p1)
//	p2 := person{"asd", 1}
//	fmt.Println(p2)
//	p3 := kjj{jjk: student{
//		"kjj",
//		"30",
//		20,
//	}, hobby: "asda"}
//	p3.jjk.name = "zzx"
//	p3.course = append(p3.course, "30")
//	fmt.Println(p3)
//	p4 := &goods{"kjj","30.000"}
//	s := reflect.TypeOf(p4).Elem()
//	fmt.Println(s)
//	for i:=0 ; i < s.NumField() ; i++ {
//		fmt.Println(s.Field(i).Tag)
//	}
//	f := funcs{30.20}
//	fmt.Println(f.getArea())
//}


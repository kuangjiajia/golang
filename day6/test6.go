package main


import (
	"fmt"
)
type person struct {
	name string
	age int
}

func (p person) printName() {
	fmt.Println(p.name)
}

func (p person) printAge() {
	fmt.Println(p.age)
}

type people interface {
	printName()
	printAge()
}

type diffPeople interface {
	printName()
	printHaha()
}

func echoArray(n interface{}) {
	fmt.Println(n)
	b,_ := n.([]int)
	fmt.Println(b)
	for i,v := range b {
		fmt.Println(i,v)
	}
}
//
//func main() {
//	var p person
//	p.name = "kuangjiajia"
//	p.printName()
//	p.printAge()
//
//	var p1 people
//	p1 = person{"zzx", 30}
//	p1.printAge()
//
//	var p2 people
//	p2 = person{"zzx",30}
//	p2.printAge()
//
//	arr := []int{1,2,3,4}
//	echoArray(arr)
//}

package main


import (
	"fmt"
)


// type person struct {
// 	name string
// 	age int
// }

// type Human struct {
// 	name string
// 	age int
// 	weight float64
// }

// type Student struct {
// 	Human
// 	hobby string
// }

// type Rectangle struct {
// 	width,height int
// }

// func area(r Rectangle) int {
// 	return r.width * r.height
// }


// func (r Rectangle) area() int {
// 	return r.width * r.height
// }

// type Box struct {
// 	color string
// 	size int
// }


// func (b *Box) setColor(c string) {
// 	b.color = c
// }

// type Human struct {
// 	name string
// 	age int
// 	birth int
// }

// func (h Human) birthday() int {
// 	return h.age * h.birth
// }

// func (h *Human) sayhi() {
// 	fmt.Printf("hello %s",h.name)
// }


type Human struct {
	name string
	age int
}

type Student struct {
	Human 
	stuId int
}

type Person struct {
	Human
	sex string
}

func (h Human) sayhi() {
	fmt.Print("hi")
}

func (h Human) sayAge() {
	fmt.Print("this age")
}

func (h Person) sayhi() {
	fmt.Print("hello")	
}

func (h Student) sayAge() {
	fmt.Print("age")		
}

type Men interface {
	sayhi()
	sayAge()
}

var a interface{} 

// 空接口可以储存任意类型的数值

// 骚操作

// var item interface{}
// type []List item List可以存任意类型的变量


func main() {
	// var p person
	// kjj := Student{Human{"Mark", 25, 120}, "Computer Science"}
	// fmt.Printf("name: %s",kjj.name)
	// s := Rectangle{10,20}
	// fmt.Print(area(s))
	// fmt.Print(s.area())
	// var b Box
	// b.setColor("red")
	// fmt.Print(b.color)
	// h := Human{"kjj",10,2010}
	// fmt.Print(h.birthday())
	// h.sayhi()
	s := Student{Human{"zzx",20},20}
	p := Person{Human{"kjj",20},"man"}
	p.sayhi()
	s.sayhi()
}
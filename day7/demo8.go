package main

import "fmt"

type Human struct {
	name string
	age int
}

type Student struct {
	Human
	hobby string
}

type Teacher struct {
	Human
	className string
}

func (h *Human) sayHi() {
	fmt.Println("hi")
}


func main() {
	kjj := Student{Human{"kjj",20},"sleep"}
	lxq := Teacher{Human{"lxq",30},"teach"}
	kjj.sayHi()
	lxq.sayHi()
}
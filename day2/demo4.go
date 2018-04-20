package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	// var p person
	// p.name = "kjj"
	// p.age = 20
	p := person{age:30,name: "kjj"}
	bob := person{age:25, name:"Bob"}

	fmt.Printf("personName: %s \npersonAge: %d \n", bob.name,bob.age)	
	fmt.Printf("personName: %s \npersonAge: %d \n", p.name,p.age)	
}
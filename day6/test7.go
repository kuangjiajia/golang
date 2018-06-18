package main

import (
	"fmt"
	"time"
)
var c chan int

func ready(w string,sec int) {
	time.Sleep(time.Duration(sec)*time.Second)
	fmt.Println(2,"is ready")
	c<-1
}

var a string
var c1 = make(chan int,10)

func f() {
	a = "hello world"
	c1<-0
}

var a2 string
var c2 = make(chan int, 10)

func f2() {
	a2 = "hello my boy"
	<-c2
}

func main() {
	c := make(chan int)

	go ready("Tee",2)
	go ready("coffee",1)
	fmt.Println("I'm waiting,but not too long")
	<-c
	<-c
	go f()
	<-c1
	fmt.Println(a)
	go f2()
	c2<-0
	fmt.Println(a2)
}
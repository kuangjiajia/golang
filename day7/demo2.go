//package main
//
//import "fmt"
//
//func main() {
//	var channel chan string = make(chan string)
//	//channel2 := make(chan int)
//	go func(message string) {
//		channel <- message
//	}("kuangjiajia")
//	fmt.Println(<- channel)
//}


package main

import (
	"fmt"
	"time"
)

type kuangjiajia struct {
	name string
	age int
}

func worker(done chan bool) {
	fmt.Println("working....")
	for i := 0 ; i < 10 ; i++ {
		time.Sleep(time.Second)
		fmt.Print(".")
	}
	kjjj := &kuangjiajia{"kuanjgiajia",20}
	fmt.Println(kjjj)
	fmt.Println("done")
	done <- true
}


func main() {
	c := make(chan bool,1)
	go worker(c)
	<- c
}
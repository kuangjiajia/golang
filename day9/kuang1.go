package main

import (
	"fmt"
	"time"
	"strconv"
)

var mes = make(chan string,3)

//func demo1() {
//	mes <- "kuangjiajia1"
//	mes <- "kuangjiajia2"
//	mes <- "kuangjiajia3"
//	mes <- "kuangjiajia4"
//}
//
//func demo2() {
//	time.Sleep(2*time.Second)
//	str := <-mes
//	str += "-change"
//	mes <- str //第一次的重新进入
//	close(mes)
//}

func channelDemo1(str chan string) {
	for i:=0; i < 10; i++ {
		str <- "chanStr" + strconv.Itoa(i)
		time.Sleep(3*time.Second)
	}
}

func channelDemo2(num chan int) {
	for i:=0; i < 10; i++ {
		num <- i
		time.Sleep(2*time.Second)
	}
}



func main() {
	//go demo1()
	//go demo2()
	//time.Sleep(3*time.Second)
	//
	//for str := range mes{
	//	fmt.Println(str)
	//}

	//fmt.Println(<-mes)
	//fmt.Println(<-mes)
	//fmt.Println(<-mes)
	//fmt.Println(<-mes)
	//str := make(chan string)
	//num := make(chan int)
	//for i:=0 ; i < 10 ; i++ {
	//	go channelDemo1(str)
	//	go channelDemo2(num)
	//}
	//
	//
	//for i:=0; i < 1000 ; i++ {
	//	select {
	//		case val,ok := <-str:
	//			if(!ok) {
	//				fmt.Println("error")
	//			}
	//			fmt.Println(val)
	//		case numVal,numOk := <-num:
	//			if(!numOk) {
	//				fmt.Println("error")
	//			}
	//			fmt.Println(numVal)
	//	}
	//}

}
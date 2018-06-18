package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int,2)
	c2 := make(chan int)

	go func() {
		var n int
		select {
		//case n = <-c1:
		case n = <-c2:
		}
		fmt.Println("数据:", n)
	}()

	fmt.Println("写入")
	c1 <- 10
	c2 <- 12

	close(c1)
	close(c2)

	time.Sleep(time.Second * 1) //等待其它协程处理

}
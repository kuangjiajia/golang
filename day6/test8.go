package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)

	for i := 0; i < 5; i++ {
		go func(n int) {
			<-c //读取到数据或通道关闭时会退出阻塞
			fmt.Println("收到通知:", n)
		}(i)
	}

	fmt.Println("广播通知")
	close(c) //关闭通道, 广播通知

	time.Sleep(time.Second * 1) //等待其它协程处理

}
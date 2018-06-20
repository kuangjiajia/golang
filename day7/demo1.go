package main

import (
	"fmt"
	"time"
)

func loop() {
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(i)
	}
}

func main() {
	loop()
	go loop()
	time.Sleep(time.Second)
}

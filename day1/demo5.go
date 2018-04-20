package main

import (
	"fmt"
)

func fib(n int) int {
	x , y := 0 , 1
	for i := 0 ; i < n ; i++ {
		x , y = y , x + y
	}
	return x
}

func main() {
	fmt.Println(fib(3))
 

	// protocal:auth:host:path(pathname,port):search(?query)hash
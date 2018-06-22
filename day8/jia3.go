package main

//import "fmt"
//
//func sum(a []int,c chan int) {
//	total := 0
//	for _,v := range a {
//		total += v
//	}
//	c <- total
//}
//
//func main() {
//	a := []int{1,3,5,52,2,3,1,3}
//	c := make(chan  int)
//	go sum(a[:len(a)/2],c)
//	go sum(a[:len(a)],c)
//	x ,y := <-c,<-c
//	fmt.Println("heheh")
//	fmt.Println(x,y,x+y)
//}

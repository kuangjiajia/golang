package main

import "fmt"


func PrintStr(str string) {
	fmt.Print(str)
}

//import "fmt"

//type Sorter interface {
//	Len() int
//	Less(i,j int) bool
//	Swap(i,j int)
//}
//
//type xi []int
//type xs []string
//
//func (p xi) Len() int {
//	return len(p)
//}
//
//func (p xi) Less(i,j int) bool{
//	return p[i] > p[j]
//}
//
//func (p xi) Swap(i, j int) {
//	p[i],p[j] = p[j],p[i]
//}
//
//
//func (p xs) Len() int {
//	return len(p)
//}
//
//func (p xs) Less(i,j int) bool{
//	return p[i] > p[j]
//}
//
//func (p xs) Swap(i, j int) {
//	p[i],p[j] = p[j],p[i]
//}
//func Sort(x Sorter) {
//	for i := 0 ; i < x.Len() - 1 ; i++ {
//		for j := i + 1 ; j < x.Len() ; j++ {
//			if x.Less(i, j) {
//				x.Swap(i, j)
//			}
//		}
//	}
//}



//func main() {
//	arr := xi{1,2,3,4,12,4,2,31}
//	Sort(arr)
//	fmt.Println(arr)
//}
//

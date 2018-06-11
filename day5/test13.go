// go语言快排

package main

import "fmt"

func bubbleSort(arr []int) []int{
    for i := 0 ; i < len(arr) ; i++ {
        for j := 0; j < len(arr) ; j++ {
            if arr[i] < arr[j] {
                arr[i], arr[j] = arr[j],arr[i]
            }
        }
    }
    return arr
}

func main() {
    var arr = []int{1,3,4,2,5,6,3}
    arr = bubbleSort(arr)
    fmt.Println(arr)
}
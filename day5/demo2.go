package main

import "fmt"

func main() {
    var arr [10]int 
    arr[0] = 10
    arr[2] = 30
    fmt.Println(arr)
    a := [3]int{1,2,3}
    fmt.Println(a)
    x := [3][2]int{[2]int{1,2},[2]int{3,4},[2]int{5,6}}
    demo := [3][2]int{{1,2},{3,4},{5,6}}
    fmt.Println(x)
    fmt.Println(demo)
    kjj := []int{1,2,3,4}
    llp := kjj[:2]
    fmt.Println(llp)
    llp[0] = 10
    fmt.Println(kjj)
    kjj = append(kjj,5)
    fmt.Println(kjj)
    s0 := []int{0,0}
    s1 := append(s0,2,3)
    fmt.Println(s1)
    s1[0] = 100
    fmt.Println(s0)
    var zzx = [...]int{0,1,2,3,4,5,6,7}
    fmt.Println(zzx)
    var pp = make([]int,6)
    ll := copy(pp,zzx[4:])
    fmt.Println(ll)
    person := map[string]int {
        "name": 10,
        "age": 20,
    }
    fmt.Println(person)
    var xxx int 
    var err bool
    xxx,err = person["name"]
    fmt.Println(err)
    fmt.Println(xxx)
    xs,ok := person["age"]
    fmt.Println(xs)
    fmt.Println(ok)
}
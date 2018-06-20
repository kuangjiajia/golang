package main

import "fmt"

type person struct {
	name string
	age int
}

func changeX(x *int) {
	(*x) += 1
}

func older(p1,p2 person) (person,int){
	if(p1.age > p2.age) {
		return p1,p1.age-p2.age
	}
	return p2,p2.age-p1.age
}

func main() {
	//kjj := new(person) //返回的是一个指针类型
	//(*kjj).name = "kjj"
	//(*kjj).age = 30
	kjj := person{"kjj",100}
	llp := person{"llp",20}
	fmt.Println(kjj,llp)
	//var x int = 20
	older , diff := older(kjj,llp)
	fmt.Println(older,diff)
	fmt.Printf("older is %s,diff is %d",older.name,diff)
	//changeX(&x)
	//fmt.Println(x)
}
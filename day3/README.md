struct 

```
    type Human struct {
        name string
        age int
        weight float64
    }

    type Student struct {
        Human 
        hobby string
    }

    kjj := Student{Human{"kjj",12,23.3},"basketball"}
    //记住Human哦
```


函数
```

type Rectangle struct {
    width,height int
}

func area(r Rectangle) int {
    return r.width * r.height
} 

调用area(x)

```

方法

```
type Rectangle struct {
    width,height int
}

func (r Rectangle) area() int {
    return r.width * r.height
}

调用
s Rectangle 
s.area()


type Color byte byte是color的别名

```


```
		for index, element := range list{
			switch value := element.(type) {
				case int:
					fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
				case string:
					fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
				case Person:
					fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
				default:
					fmt.Println("list[%d] is of a different type", index)
			}
		}
```
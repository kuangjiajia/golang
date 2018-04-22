遇到const iota 默认值为0 叠加 声明的时候 
每遇到一个iota重置为0
const特有

作用：（一次声明的常量个数？）

array到slice方法
引用类型 是

append 方法  arr = append(arr,newProp)
copy 方法 copy(slice1,slice2)

map+make

var type map[propType]valueType
type = make(map[propType]valueType)
初始化
type[prop] = value

name := map[string]float32{}


value , ok := obj[prop]
ok => prop是否存在 返回布尔值
value => 返回obj[prop]的值
注意 也是引用类型


流程控制语句

if else 在if语句中可以声明变量，只在语句中起作用

goto语句 

for循环 
for expression1 ; expression2 ; expression3 {

}
前提是变量已初始化
for expression2 {

}
break 跳出循环
continue 跳出本轮循环，开始下一轮循环

range 读取map的数据

for v , i := range map {
    fmt.Print 
}

switch 

swicth caseValue {
    case tmp1:  
        xxx
    case tmp2:
        xxx
    case tmp3:
        xxx
    default:
        xxx
}
默认break 就是说不用+break了
fallthrough 强制执行

func returnFunc (input1 type1,input2 type2) (type1,type2){
    return xxx
}
func returnArg(agr,...int) {

}
arg ...int告诉Go这个函数接受不定数量的参数。注意，这些参数的类型全部是int

defer
defer 指定的函数在推出程序前调用


```
    func ReadWrite() bool {
        defer fs.close()
        if(xxx)
            return true
        else
            return false
    }
``` 


引入包
```
点操作
    import (
        . "fmt"
    )
    fmt.Print = Print
别名操作
    import (
        f "fmt"
    )
    fmt.Print = f.Print
_操作
	import (
	    "database/sql"
	    _ "github.com/ziutek/mymysql/godrv"
	)
_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
```



struct 类型

type person struct {
    name string
    age int
}

var kjj Person 


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
package main
//
//import "fmt"
//
////type obj map[string]string
//
//const (
//	WHITE = iota
//	BLACK
//	RED
//	YELLOW
//	GREEN
//)
//
//type Color byte
//
//type BOX struct {
//	width,height,depth float64
//	color Color
//}
//
//type BoxList []BOX
//
//func (b BOX) volume() float64 {
//	return b.width * b.height * b.depth
//}
//
//func (b *BOX) SetColor(c Color) {
//	b.color = c
//}
//
//func (b1 BoxList) BiggestColor() Color {
//	v := 0.00
//	k := Color(WHITE)
//	for _,b := range b1 {
//		if bv := b.volume() ; bv > v {
//			v = bv
//			k = b.color
//		}
//	}
//	return  k
//}
//
//func (b1 BoxList) PaintItBlack() {
//	for i,_ := range b1 {
//		fmt.Println(b1[i])
//		b1[i].SetColor(BLACK)
//	}
//}
//
//func change(x *int) {
//	*x += 1
//}
//
//func (c Color) String() string {
//	strings := []string{"WHITE","BLACK","RED","YELLOW","GREEN"}
//	return strings[c]
//}
//
//func main() {
//	//kjj := obj {
//	//	"name": "kjj",
//	//	"hobby": "sss",
//	//}
//	//fmt.Println(kjj)
//
//	boxes := BoxList {
//		{20.22,30.33,80,WHITE},
//		{26.22,30.33,70,BLACK},
//		{20.22,40,80,RED},
//		{26.22,30.33,80,GREEN},
//		{24.22,30.33,60,YELLOW},
//	}
//	fmt.Printf("boxes have %d in your set",len(boxes))
//	fmt.Println("the volume of the first one is",boxes[0].volume())
//	fmt.Println("the color of the last one is",boxes[len(boxes)-1].color.String())
//	fmt.Println("the biggest one is",boxes.BiggestColor().String())
//	boxes.PaintItBlack()
//	fmt.Println(boxes)
//	x := 10
//	changeX(&x)
//}

//go语言指针的特征 => 自动识别指针 函数传入指针 操作时不用+"*"
//若传入对象 调用函数的时候不用 &，自动转为地址
//注意 都是操作对象的时候
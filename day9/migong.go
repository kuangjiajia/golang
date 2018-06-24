package main

//import (
//	"os"
//	"fmt"
//)

//func readMaze(filename string) [][]int {
//	file,err := os.Open(filename)
//	if err != nil{
//		panic(err)
//	}
//	var row,col int
//	fmt.Fscanf(file,"%d %d",&row,&col) // fmt.fscanf	扫描 换行当作空处理
//	maze := make([][]int,row) //行的总数
//	for i := range maze { // i => 第几行
//		maze[i] = make([]int,col) // 每行对应的列数
//		for j := range maze[i] {
//			fmt.Fscanf(file,"%d",&maze[i][j])
//		}
//	}
//	return maze
//}
//
//type Point struct {
//	i,j int
//}
//
//var dirs = [4]Point {
//	{-1,0},
//	{0,-1},
//	{1,0},
//	{0,1},
//}
//
//func (p Point) add(r Point) Point {
//	return Point{p.i+r.i,p.j+r.j}
//}
//
//func (p Point) at(grid [][]int) (int,bool) { //判断是否越界
//	if p.i < 0 || p.i >= len(grid) {
//		return 0,false
//	}
//	if p.j < 0 || p.j >= len(grid[p.i]){
//		return 0,false
//	}
//	return grid[p.i][p.j],true
//}
//
//func walk(maze [][]int,start,end Point) [][]int{
//	// 初始化steps
//	steps := make([][]int,len(maze))
//	for i := range steps {
//		steps[i] = make([]int,len(maze[i]))
//	}
//	// 初始化Q
//	Q := []Point{start}
//	fmt.Println(Q)
//
//	for len(Q) > 0 {
//		cur := Q[0]
//		Q = Q[1:]
//
//		if cur == end {
//			break
//		}
//		for _,dir := range dirs {
//			//next := cur+dir
//			next := cur.add(dir)
//			val,ok := next.at(maze)
//			if !ok || val == 1 {
//				continue
//			}
//			val,ok = next.at(steps)
//			if !ok || val != 0 { //周围有东西 但是为1
//				continue
//			}
//			if next == start {
//				continue
//			}
//
//			curSteps,_ := cur.at(steps)
//			steps[next.i][next.j] = curSteps+1
//			Q = append(Q,next)
//		}
//	}
//	return steps
//}
//
//func main() {			fmt.Println(curSteps)
//
//	maze := readMaze("./data.in")
//	for _,row := range maze {
//		for _,val := range row {
//			fmt.Printf("%d ",val )
//		}
//		fmt.Println("")
//	}
//	fmt.Println(len(maze))
//	steps := walk(maze,Point{0,0},Point{len(maze)-1,len(maze[0])-1})
//	for _,row := range steps {
//		for _,val :=range row {
//			fmt.Printf("%3d",val)
//		}
//		fmt.Println("")
//	}
//}
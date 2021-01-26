package datastructure

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y int
}

type Vertex2 struct {
	X int
	Y float64
	Z string
}

var (
	v1 = Vertex{3, 4}
	v2 = Vertex{X: 8}
	v3 = Vertex{}
	p  = &Vertex{5,6}
	v4 = Vertex2{}
	v5 = Vertex2{Y: 2.8, Z: "Liusy"}
)

func TestPointer() {
	fmt.Println("Enter TestPointer......................")
	i, j := 42, 81
	
	p := &i
	//fmt.Println(&i)   // &取地址即指针，*取值。
	fmt.Println(p)   	// 指针变量p指向的内存地址
	fmt.Println(&p)  	// &p就表示编译器为指针变量p分配的内存地址
	fmt.Println(*p)  	// *p表示此指针指向的内存地址中存放的内容
	fmt.Println(*p/2)
	*p = 88
	fmt.Println(i)

	p = &j
	fmt.Println(*p)
	fmt.Println(math.Sqrt(float64(*p)))
	*p *= *p
	fmt.Println(j)
	fmt.Println("Leave TestPointer......................")
	fmt.Println()
}

func TestStruct(){
	v := Vertex{1, 2}
	p := &v
	p.X = 186
	fmt.Println(v)
}

func TestStruct2(){
	fmt.Println(v1, v2, v3, p, *p, v4, v5)
}
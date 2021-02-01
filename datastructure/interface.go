package datastructure

import (
	"fmt"
	"math"
)

/* 定义接口 */
type Abser interface{
	Abs() float64
}

type MyfloatI float64

/* 定义结构体 */
type VertexI struct {
	X, Y float64
}

/* 实现接口方法 */
func (f MyfloatI) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

/* 实现接口方法 */
func (v *VertexI) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestInterface(){
	fmt.Println("Enter TestInterface......................")
	var i Abser
	f := MyfloatI(-math.Sqrt2)
	v := VertexI{3, 4}
	
	i = f
	fmt.Println(i, i.Abs())
	
	i = &v
	fmt.Println(i, i.Abs())
	fmt.Println("Leave TestInterface......................")
	fmt.Print()
}
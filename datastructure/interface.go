package datastructure

import (
	"fmt"
	"math"
)

type Abser interface{
	Abs() float64
}

type MyfloatI float64

type VertexI struct {
	X, Y float64
}

func (f MyfloatI) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

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
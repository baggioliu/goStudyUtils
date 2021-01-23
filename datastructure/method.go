package datastructure

import (
	"fmt"
	"math"
)

type Myfloat float64

type VertexM struct {
	X, Y float64
}

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func (v VertexM) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//use pointer receiver gain better performance
func (v *VertexM) AbsP() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *VertexM) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func TestMethod(){
	v := &VertexM{5, 12}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(10)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
	
	f := Myfloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	
	p := &VertexM{6, 8}
	fmt.Println(p.AbsP())
}
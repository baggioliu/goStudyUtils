package datastructure

import (
	"fmt"
	"math"
	"strings"
)

type Mytype struct {
	Int int
	Str string
}

//nested struct
type Location struct {
	Lat, Long float64
	Mytype
}

var m map[string]Location

var m2 = map[string]Location{
	"Bell Labs": {
		40.68433, -74.39967, Mytype{1, "liuyu"},
	},
	"Google": {
		37.42202, -122.08408, Mytype{2, "liusy"},
	},
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func TestMap() {
	fmt.Println("Enter TestMap......................")
	m = make(map[string]Location)
	m["Beijing"] = Location{116.40, 39.90, Mytype{3, "Liu"}}
	fmt.Println(m["Beijing"])
	fmt.Println(m)
	fmt.Println(m2)
	
	for key,value := range m2 {
		fmt.Println(key, value)
	}
	
	m3 := make(map[string]int)
	m3["Score"] = 100
	fmt.Println("The value:", m3["Score"])
	
	m3["Score"] = 120
	fmt.Println("The value:", m3["Score"])
	
	v,ok := m3["Score"]
	fmt.Println("The value:", v, "Present?", ok)
	
	delete(m3, "Score")
	fmt.Println("The value:", m3["Score"])
	
	v2,ok2 := m3["Score"]
	fmt.Println("The value:", v2, "Present?", ok2)
	
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	
	pos, neg := add(), add()
	for i:=0; i<10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	
	wcs := "x y z x y z x x x x y z z x total bin test text bin total" 
	fmt.Println(wcs)
	fmt.Println(WordCount(wcs))
	fmt.Println("Leave TestMap......................")
	fmt.Println()
}

func add() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 返回一个“返回int的函数”
func fibonacci() func() int {
	back1, back2 := 0, 1
	return func() int {
		back1, back2 = back2, back1 + back2
		return back1
	}
}

func WordCount(s string) map[string]int {
	//var arrays []string = strings.Fields(s)
	arrays := strings.Fields(s)
	m := make(map[string]int)
	for i:=0; i<len(arrays); i++ {
/*
		v,ok := m[arrays[i]]
		if !ok {
			m[arrays[i]] = 1
		} else {
			m[arrays[i]] = v + 1
		}
*/
		v := m[arrays[i]]
		m[arrays[i]] = v + 1
	}
	return m
}
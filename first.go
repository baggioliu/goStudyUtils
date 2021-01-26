package main

import (
	"fmt"
	"github.com/baggioliu/goStudyUtils/datastructure"
	"math"
	"math/rand"
	"runtime"
	"time"
)

const (
	Son string = "Liusy"
	FaviNumber int = 8
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

var (
	i, j int = 3, 4
    c, java, python, _go bool = false, false, false, true
	s = "Hello " + Son + "!"
    k float64 = math.Pi
	cpx = 5i
	big uint64 = 1<<64 - 1
)

func sum(x, y int, z float64) (result float64) {
	result = float64(x) + float64(y) + z
	return
}

func needInt(x int) int { 
	return x*10 + 1 
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了, scope of v only belongs to if and else block
	return lim
}

func countdown() {
	//defer 栈
    //推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	fmt.Println("Enter main func......................")
	fmt.Println(c, java, python, _go, s)
	fmt.Printf("Type: %T Value: %v\n", j, j)
	fmt.Printf("Type: %T Value: %v\n", _go, _go)
	fmt.Printf("Type: %T Value: %v\n", s, s)
	fmt.Printf("Type: %T Value: %v\n", k, k)
	fmt.Printf("Type: %T Value: %v\n", cpx, cpx)
	
	ni1 := needInt(Small)
	nf1 := needFloat(Small)
	nf2 := needFloat(Big)
	fmt.Printf("Type: %T Value: %v\n", ni1, ni1)
	fmt.Printf("Type: %T Value: %v\n", nf1, nf1)
	fmt.Printf("Type: %T Value: %v\n", nf2, nf2)
	
	for m := 1; m < FaviNumber; m++ {
		fmt.Println(sum(i+j, rand.Intn(m), k))
	}
	
	var f float64 = math.Sqrt(math.Pow(float64(i), 2) + math.Pow(float64(j), 2))
	var z uint = uint(f)
	fmt.Println(i, j, f, z, big)
	
	fmt.Println(pow(3,2,10), pow(3,3,20))
	
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, // plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(int(today))
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	
	countdown()
	
	fmt.Println()

	datastructure.TestPointer()
	datastructure.TestStruct()
	datastructure.TestStruct2()

	fmt.Println()

	v := datastructure.Vertex{X: 1, Y: 2}
	p := &v
	fmt.Println(p) 
	fmt.Println(*p)
	p.X = 138
	fmt.Println(v)
	
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var slices []int = primes[1:4]
	fmt.Println(slices)
	
	s1 := primes[0:4]
	s2 := primes[2:6]
	fmt.Println(s1, s2)
	s1[2] = 888
	s2[1] = 999
	fmt.Println(s1, s2)  //对切片的修改，会影响到底层数组
	fmt.Println(primes)

	fmt.Println()

	datastructure.TestSlice()
	datastructure.TestMap()
	datastructure.TestMethod()
	datastructure.TestInterface()

	fmt.Println("Quit main func......................")

}
package datastructure

import (
	"fmt"
	"strconv"
	"strings"
)

func TestSlice() {
	fmt.Println("Enter TestSlice......................")
	q := []int{2, 3, 5, 7, 11, 13, 17, 19}
	fmt.Println(q)
	
	for i,v := range q {
		fmt.Printf("Index is %d, value is %d\n", i, v)
	}

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	
	q = q[:0]
	printSlice(q)
	
	q = q[:6]
	printSlice(q)
	
	q = q[3:]
	printSlice(q)

	// make([]T, length, capacity) 这里 length 是数组的长度并且也是切片的初始长度。
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
	
	//slice of slice:
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	
	
	var s1 []int
	printSlice(s1)

	// 添加一个空切片
	s1 = append(s1, 0)
	printSlice(s1)

	// 这个切片会按需增长
	s1 = append(s1, 1)
	printSlice(s1)

	// 可以一次性添加多个元素
	s1 = append(s1, 2, 3, 4)
	printSlice(s1)
	
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << i  // 2**i
	}
	for _,value := range pow {
		fmt.Printf("%d\n", value)
	}
	
	fmt.Println(strconv.ParseInt("123", 10, 32))
	
	fmt.Println(Pic(5, 3))
	fmt.Println("Leave TestSlice......................")
	fmt.Println()
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, s=%v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	s1 := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s2 := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			s2[j] = uint8(i+j)
			//fmt.Println(s2)
		}
		s1[i] = s2
	}
	return s1
}
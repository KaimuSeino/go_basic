package main

import (
	"fmt"
	"unsafe"
)

const secret = "abc"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

// 一括に値を定義している
var (
	i int
	s string
	b bool
)

func main() {
	// var i int
	// var i int = 2
	// i := 1
	// ui := uint16(2)
	// fmt.Println(i)
	// fmt.Printf("i: %v %T\n", i, i)
	// fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	// f := 1.23456789
	// s := "hello"
	// b := true
	// pi, title := 3.14, "GO"
	// fmt.Printf("f %[1]v %[1]T\n", pi)
	// fmt.Printf("s %[1]v %[1]T\n", title)
	// fmt.Printf("b %[1]v %[1]T\n", b)

	// 型変換
	// x := 10
	// y := 1.23
	// z := float64(x) + y
	// fmt.Println(z)

	// fmt.Printf("Mac: %v Windows: %v Linux: %v\n", Mac, Windows, Linux)
	// i := 1
	// i = 2
	// fmt.Printf("i: %v\n", i)
	// i += 1
	// fmt.Printf("i: %v\n", i)

	// ポインタ
	var ui1 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1)
	var ui2 uint16
	fmt.Printf("memory address of ui2: %p\n", &ui2)

	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)

	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("value of ui1(dereference): %v\n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))
	fmt.Printf("value of p1(dereference): %v\n", *pp1)
	fmt.Printf("value of ui1(dereference): %v\n", **pp1)
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n", ui1)

	// shadowing
	ok, result := true, "A"
	fmt.Printf("memory address of result: %p\n", &result)
	if ok {
		result = "B" // :=だとメモリアドレスが違う場所に保存される。
		fmt.Printf("memory address of result: %p\n", &result)
		println(result)
	} else {
		result = "C"
		println(result)
	}
	println(result)
}

// fmt.Println(calculator.Offset)
// fmt.Println(calculator.Sum(2, 3))
// fmt.Println("Multiply: ", calculator.Multiply(3, 4))

// fmt.Println("hello world")
// sl := []int{1, 2, 3}
// if len(sl) > 0 {
// 	fmt.Println("unreachable code")
// }

// godotenv.Load()
// fmt.Println(os.Getenv("GO_ENV"))

package main

import (
	pointershadowing "go-basics/pointer-shadowing"
)

const secret = "abc"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

// 一括に値を定義している
// var (
// 	i int
// 	s string
// 	b bool
// )

func main() {
	// pointershadowing.Pointer()
	pointershadowing.Shadowing()
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

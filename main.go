package main

import (
	"fmt"
	"go-basics/calculator"
)

func main() {
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(2, 3))
	fmt.Println("Multiply: ", calculator.Multiply(3, 4))
}

// fmt.Println("hello world")
// sl := []int{1, 2, 3}
// if len(sl) > 0 {
// 	fmt.Println("unreachable code")
// }

// godotenv.Load()
// fmt.Println(os.Getenv("GO_ENV"))

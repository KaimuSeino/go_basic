package pointershadowing

import "fmt"

func Shadowing() {
	ok, result := true, "A"
	fmt.Printf("memory address of result: %p\n", &result)
	if ok {
		result := "B"
		fmt.Printf("memory address of result; %p\n", &result)
		println(result)
	} else {
		result := "C"
		println(result)
	}

	println(result)
}

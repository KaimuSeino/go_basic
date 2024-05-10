package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type customConstraints interface {
	~int | int16 | float32 | float64 | string
}

type NewInt int // この型をcustomConstraintsに入れたいのでintに~をつけています↑

func add[T customConstraints](x, y T) T {
	return x + y
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func sumValues[K int | string, V constraints.Float | constraints.Integer](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func Generics() {
	fmt.Printf("%v\n", add(1, 6))
	fmt.Printf("%v\n", add(1.2, 4.3))
	fmt.Printf("%v\n", add("file", ".csv"))
	var i1, i2 NewInt = 3, 4
	fmt.Printf("%v\n", add(i1, i2))
	fmt.Printf("%v\n", min(i1, i2))

	m1 := map[string]uint{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	m2 := map[int]float32{
		1: 1.23,
		2: 2.34,
		3: 3.45,
	}
	fmt.Println(sumValues(m1))
	fmt.Println(sumValues(m2))
}

package variables

import (
	"fmt"
	"unsafe"
)

func Varitables() {
	// var i int

	// p := &i

	// fmt.Println(&i)
	// fmt.Printf("i: %v[bytes]\n", unsafe.Sizeof(i))
	// fmt.Println(p)
	// fmt.Printf("p: %v[bytes]\n", unsafe.Sizeof(p))

	i := 1
	p := &i
	fmt.Printf("i: %v[bytes]\n", unsafe.Sizeof(i))
	fmt.Println(p)

}

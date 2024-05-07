package slicemap

import "fmt"

func Slice() {
	// 配列の定義
	// goの配列は最初に宣言した配列のサイズを変えることはできない。
	var a1 [3]int
	var a2 = [3]int{10, 20, 30}
	a3 := [...]int{10, 20}
	fmt.Printf("%v %v %v\n", a1, a2, a3)
	fmt.Printf("%v %v\n", len(a3), cap(a3))
	fmt.Printf("%T %T\n", a2, a3)

	// sliceの定義
	var s1 []int
	s2 := []int{}
	fmt.Printf("s1: type: %[1]T value: %[1]v length: %v capacity: %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: type: %[1]T value: %[1]v length: %v capacity: %v\n", s2, len(s2), cap(s2))
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1: type: %[1]T value: %[1]v length: %v capacity: %v\n", s1, len(s1), cap(s1))
	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	fmt.Printf("s1: type: %[1]T value: %[1]v length: %v capacity: %v\n", s1, len(s1), cap(s1))

	s4 := make([]int, 0, 2)
	fmt.Printf("s4: type: %[1]T value: %[1]v length: %v capacity: %v\n", s4, len(s4), cap(s4))
	s4 = append(s4, 1, 2, 3, 4)
	fmt.Printf("s4: type: %[1]T value: %[1]v length: %v capacity: %v\n", s4, len(s4), cap(s4))

	s5 := make([]int, 4, 6)
	fmt.Printf("s5: type: %[1]T value: %[1]v length: %v capacity: %v\n", s5, len(s5), cap(s5))
	s6 := s5[1:3]
	s6[1] = 10
	fmt.Printf("s5: type: %[1]T value: %[1]v length: %v capacity: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: type: %[1]T value: %[1]v length: %v capacity: %v\n", s6, len(s6), cap(s6))
	s6 = append(s6, 2)
	fmt.Printf("s5: type: %[1]T value: %[1]v length: %v capacity: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6 appended: type: %[1]T value: %[1]v length: %v capacity: %v\n", s6, len(s6), cap(s6))

	sc6 := make([]int, len(s5[1:3]))
	fmt.Printf("s5 source of copy: value: %v length: %v capacity: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("sc6 dst copy defore: value: %v length %v capacity: %v\n", sc6, len(sc6), cap(sc6))
	copy(sc6, s5[1:5])
	fmt.Printf("sc6 dst of copy after: value: %v length %v capacity: %v\n", sc6, len(sc6), cap(sc6))
	sc6[1] = 12
	fmt.Printf("s5: value: %v length: %v capacity: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("sc6: value: %v length: %v capacity: %v\n", sc6, len(sc6), cap(sc6))

	s5 = make([]int, 4, 6)
	fs6 := s5[1:3:3] //indexが2までメモリを共有する
	fmt.Printf("s5: value: %v length: %v capacity: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6: value: %v length: %v capacity: %v\n", fs6, len(fs6), cap(fs6))

	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8)
	// この出力はindexが２までは反映されるが、それ以外は反映されないことを表している。
	fmt.Printf("s5: value: %v length: %v capacity: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6: value: %v length: %v capacity: %v\n", fs6, len(fs6), cap(fs6))

	s5[3] = 9
	fmt.Printf("s5: value: %v length: %v capacity: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6: value: %v length: %v capacity: %v\n", fs6, len(fs6), cap(fs6))
}

package slicemap

import "fmt"

func Map() {
	// map[キー]値
	var m1 map[string]int
	m2 := map[string]int{}
	fmt.Printf("m1: value: %v is_nil: %v\n", m1, m1 == nil)
	fmt.Printf("m1: value: %v is_nil: %v\n", m2, m2 == nil)

	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	fmt.Printf("m2: value: %v length %v m2['A']: %v\n", m2, len(m2), m2["A"])
	delete(m2, "A")
	fmt.Printf("m2: value: %v length %v m2['A']: %v\n", m2, len(m2), m2["A"])
	v, ok := m2["A"] // 第二引数に値が存在するかしないかを決める
	fmt.Printf("m2: value: %v is_ok: %v\n", v, ok)
	v, ok = m2["C"]
	fmt.Printf("m2: value: %v is_ok: %v\n", v, ok)

	for k, v := range m2 {
		fmt.Printf("m2: key: %v value: %v\n", k, v)
	}
}

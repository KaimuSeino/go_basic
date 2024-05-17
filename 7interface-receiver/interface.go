package interfacereceiver

import (
	"fmt"
	"unsafe"
)

// interfaceではメソッドの一覧を定義できる
type controller interface {
	speedUp() int // 引数なし 返り値がint
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}
type bycycle struct {
	speed      int
	humanPower int
}

// vehicleのポインタを引数として受け取るreceiver
func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}
func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

// bycycleのポインタを引数として受け取るreceiver
func (b *bycycle) speedUp() int {
	b.speed += 3 * b.humanPower
	return b.speed
}
func (b *bycycle) speedDown() int {
	b.speed -= 1 * b.humanPower
	return b.speed
}

func speedUpAndDown(c controller) {
	fmt.Printf("current Up speed: %v\n", c.speedUp())
	fmt.Printf("current Down speed: %v\n", c.speedDown())
}

// stringer interfaceを使ってみる
func (v vehicle) String() string {
	// Sprintfは標準出力ではなくstringで返す
	return fmt.Sprintf("Vehicle current speed is %v (enginePower %v)", v.speed, v.enginePower)
}

func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

// main
func Interface() {
	v := &vehicle{
		speed:       0,
		enginePower: 5,
	}
	speedUpAndDown(v)
	b := &bycycle{
		speed:      0,
		humanPower: 3,
	}
	speedUpAndDown(b)

	// goに標準で実装されているstringer interfaceを使ってみる。
	// Println内でString interfaceを実装しているかのチェックを行います。
	// Stringer interfaceを持つ場合は必ず、StringのメソッドString()が存在するので、そのString()を実行する
	fmt.Println(v)

	// 全ての型を取ることができるany型
	var i1 interface{}
	var i2 any
	fmt.Printf("i1 value: %[1]v type: %[1]v size: %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("i2 value: %[1]v type: %[1]v size: %v\n", i2, unsafe.Sizeof(i2))
	checkType(i2)
	i2 = 2
	checkType(i2)
	i2 = "string"
	checkType(i2)
	fmt.Printf("i2 value: %[1]v type: %[1]v size: %v\n", i2, unsafe.Sizeof(i2))
}

package calculator

import "fmt"

var offset float64 = 1 // パッケージ内のみで使えるプライベートな変数
var Offset float64 = 1 // 外部のパッケージでも使用できる変数

func Sum(a float64, b float64) float64 { // 外部のパッケージでも使用できる関数
	fmt.Println("multiply: ", multiply(a, b))
	return a + b + offset
}

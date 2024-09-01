package dependencyinjection

/*
依存関係の基本概念
「クラスが他のクラスの機能を利用するために必要とするオブジェクト」、つまり
「依存関係」とは、例えば車とエンジンの関係のようなものである。
車が動くためにはエンジンが必要。この時、車はエンジンに依存していると言える。
*/

import "fmt"

// エンジンの構造体
type Engine struct{}

// エンジンに発車のメソッドを追加する
func (e *Engine) Start() {
	fmt.Println("Engine started")
}

// Carの構造体
type Car struct {
	engine *Engine
}

// 車を作る関数
func NewCar() *Car {
	return &Car{
		engine: &Engine{}, // 車が自分でエンジンを作成している
	}
}

// 車にドライブのメソッドを追加
func (c *Car) Drive() {
	c.engine.Start()
	fmt.Println("Car is driving")
}

func LetsDrive() {
	car := NewCar()
	car.Drive()
}

// 上記はCarクラスがEngineクラスに依存している

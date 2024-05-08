package controlflow

import (
	"fmt"
	"time"
)

type item struct {
	price float32
}

func ControlFlow() {
	a := -1
	if a == 0 {
		fmt.Println("ゼロ！！")
	} else if a > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for {
	// 	fmt.Println("working")
	// 	time.Sleep(2 * time.Second)
	// }

	var i int
	for {
		if i > 3 {
			break
		}
		fmt.Println(i)
		i += 1
		time.Sleep(1000 * time.Millisecond)
	}

loop: // for文に名前をつけている
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			time.Sleep(1000 * time.Millisecond)
			continue
		case 7:
			time.Sleep(1000 * time.Millisecond)
			continue
		case 9:
			time.Sleep(1000 * time.Millisecond)
			break loop
		default:
			fmt.Printf("%v", i)
		}
	}
	fmt.Printf("\n")

	items := []item{
		{price: 10.0},
		{price: 20.0},
		{price: 30.0},
	}

	// これでは構造体のコピーが生成されただけなので、itemsの値は変化しない
	for _, i := range items {
		i.price *= 1.1
	}
	fmt.Printf("items: %+v\n", items)

	// indexを指定して直接itemsの値を変更する
	for i := range items {
		items[i].price *= 1.1
	}
	fmt.Printf("items: %+v\n", items)
}

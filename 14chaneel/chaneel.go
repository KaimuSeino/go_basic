package chaneel

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Chaneel() {
	// make(chan)でchaneelを作成
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(time.Second)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	// fmt.Println("Chaneel function finish")

	// goroutine leak: ずっと稼働し続けてメモリが解放されないgoroutineが存在すること
	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	fmt.Printf("number of working goroutines: %d\n", runtime.NumGoroutine())

	// buffer付きのchaneel
	ch2 := make(chan int, 1) // 第二引数にbufferのサイズを指定
	ch2 <- 2
	fmt.Println(<-ch2)
	// bufferがいっぱいになるまでは、chaneelの受信(<-ch2)を待たなくても書き込み(ch2 <- 2)ができる
	// 書き込みと読み込みの順番を逆にする→deadlockが発生！
	// fmt.Println(<-ch2) // ch2の読み込みができない
	// ch2 <- 2

	// 次にch2に書き込みを2回行う→deadlock発生
	// ch2 <- 2 // bufferが１だから書き込み可能！
	// ch2 <- 3 // 無理
	// fmt.Println(<-ch2)

	// channel close
	ch3 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch3)
	}()
	ch3 <- 10
	close(ch3)
	v, ok := <-ch3
	fmt.Printf("%v %v\n", v, ok)
	wg.Wait()

	ch4 := make(chan int, 2)
	ch4 <- 2
	ch4 <- 3
	close(ch4)
	v, ok = <-ch4 // ch4の1つ目の値を取り出す
	fmt.Printf("%v %v\n", v, ok)
	v, ok = <-ch4 // ch4の2つ目の値を取り出す
	fmt.Printf("%v %v\n", v, ok)
	v, ok = <-ch4
	fmt.Printf("%v %v\n", v, ok)
	// bufferのchaneelがcloseしても、まだ読み込まれていない場合は、それを読み込んだ後にflagの値がfalseになる

	// capsel化
	ch5 := generateCountStream()
	for v := range ch5 {
		fmt.Println(v)
		time.Sleep(2 * time.Second)
	}

	// notification
	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v started\n", i)
			<-nCh // ここで受信待ちをしていてdeadlock!
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(nCh)
	// chaneelをcloseすると、受信待ちのブロッキングが一斉に解除されて残りの表示の処理を進める
	fmt.Println("unblocked by manual close")

	wg.Wait()
	fmt.Println("finish")
}

// 関数の返り値に<-chan intを渡すことで、この関数は読み込み専用の関数として扱える。
func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
			fmt.Println("write")
		}
	}()
	return ch
}

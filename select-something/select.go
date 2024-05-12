package selectsomething

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const bufSize = 5

func Select() {
	// with timeout context
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		ch1 <- "A"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		ch2 <- "B"
	}()

loop:
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done(): // ctxのDoneChaneelの読み込み→タイムアウトが発生した時にfor文を抜ける
			fmt.Println(ctx.Err())
			break loop
		case v := <-ch1: // ch1に書き込みの値がある場合、または、ch1がcloseされた場合に実行
			fmt.Println(v)
			ch1 = nil
		case v := <-ch2:
			fmt.Println(v)
			ch2 = nil
		}
	}
	wg.Wait()
	fmt.Println("finish")

	// default case
	ch3 := make(chan string, bufSize)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < bufSize; i++ {
			time.Sleep(1000 * time.Millisecond)
			ch3 <- "hello"
		}
	}()
	for i := 0; i < bufSize; i++ {
		select {
		case m := <-ch3:
			fmt.Println(m)
		default:
			fmt.Println("no msg arrived")
		}
		time.Sleep(1500 * time.Millisecond)
	}

	// receive continuous data
	ch4 := make(chan int, bufSize)
	ch5 := make(chan int, bufSize)
	ctx1, cancel1 := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel1()

	wg.Add(3)
	go countProducer(&wg, ch4, bufSize, 50)
	go countProducer(&wg, ch5, bufSize, 500)
	go countConsumer(ctx1, &wg, ch4, ch5)
	wg.Wait()
	fmt.Println("finish")
}

func countProducer(wg *sync.WaitGroup, ch chan<- int, size int, sleep int) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < size; i++ {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ch <- i
	}
}

func countConsumer(ctx context.Context, wg *sync.WaitGroup, ch1 <-chan int, ch2 <-chan int) {
	defer wg.Done()

loop:
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break loop
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Printf("ch1 %v\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Printf("ch2 %v\n", v)
		}
	}
}

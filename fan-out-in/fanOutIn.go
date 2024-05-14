package fanoutin

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func FanOutIn() {
	cores := runtime.NumCPU()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	outChs := make([]<-chan string, cores)
	inData := generator(ctx, nums...)
	for i := 0; i < cores; i++ {
		// fanOut関数の処理をCPUの数だけ分散させる
		outChs[i] = fanOut(ctx, inData, i+1)
	}
	var i int
	flag := true
	for v := range fanIn(ctx, outChs...) {
		if i == 4 {
			cancel()
			flag = false
		}
		if flag {
			fmt.Println(v)
		}
		i++
	}
	fmt.Println("finish")
}

func generator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

// pipelineを流れてくるchの値を複数のgoroutineに分散させて並行処理
func fanOut(ctx context.Context, in <-chan int, id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		// 大きい処理を擬似的に追加
		heavyWork := func(i int, id int) string {
			time.Sleep(200 * time.Millisecond)
			return fmt.Sprintf("result: %v (id: %v)", i*i, id)
		}
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- heavyWork(v, id):
			}
		}
	}()
	return out
}

func fanIn(ctx context.Context, chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	multiplex := func(ch <-chan string) {
		defer wg.Done()
		for text := range ch {
			select {
			case <-ctx.Done():
				return
			case out <- text:
			}
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go multiplex(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

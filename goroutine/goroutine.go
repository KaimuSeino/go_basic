package goroutine

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func Goroutine() {
	// 関数の先頭にgoをつけることでgoroutineになる
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done() // ここでwgが-1される
	// 	fmt.Println("goroutine invoked")
	// }()
	// wg.Wait()
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	// fmt.Println("main func finish")

	// tracer
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error", err)
		}
	}()
	// go tool trace trace.out
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error", err)
	}
	defer trace.Stop()
	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

	// task(ctx, "task1")
	// task(ctx, "task2")
	// task(ctx, "task3")
	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "Task1")
	go cTask(ctx, &wg, "Task2")
	go cTask(ctx, &wg, "Task3")
	wg.Wait()

	// goroutineを使う時の注意
	s := []int{1, 2, 3}
	for _, i := range s {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main func finish")
}

func task(ctx context.Context, name string) {
	defer trace.StartRegion(ctx, name).End()
	time.Sleep(time.Second)
	fmt.Println(name)
}

func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}

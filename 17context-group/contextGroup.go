package contextgroup

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextGroup() {
	// context: 一番の用途はmain goroutineからsub goroutineを一斉にキャンセルすること
	// timeout
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	wg.Add(3)
	go subTask01(ctx, &wg, "A")
	go subTask01(ctx, &wg, "B")
	go subTask01(ctx, &wg, "C")
	wg.Wait()

	fmt.Println("#####cancel#####")
	// cancel
	// criticalTask関数がタイムアウトしたらnormal関数の処理も終了する。
	ctx, cancel = context.WithCancel(context.Background())
	// WithCancelのcancel関数が実行されるとctx.Done()がtrueになるため
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		v, err := criticalTask(ctx)
		if err != nil {
			fmt.Printf("critical task cancelled due to: %v\n", err)
			cancel()
			return
		}
		fmt.Println("success", v)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		v, err := normalTask(ctx)
		if err != nil {
			fmt.Printf("normal task cancelled due to: %v\n", err)
			return
		}
		fmt.Println("success", v)
	}()
	wg.Wait()

	fmt.Println("#####deadline#####")
	// deadline
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(20*time.Millisecond))
	defer cancel()
	ch := subTask02(ctx)
	v, ok := <-ch
	if ok {
		fmt.Println(v)
	}
	fmt.Println("finish")

}

func subTask02(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		deadline, ok := ctx.Deadline()
		if ok {
			if deadline.Sub(time.Now().Add(30*time.Millisecond)) < 0 {
				fmt.Println("impossible to meet deadline")
				return
			}
		}
		time.Sleep(30 * time.Millisecond) // task02が30msある
		ch <- "hello"
	}()
	return ch
}

// cancel
func criticalTask(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
	defer cancel()
	t := time.NewTicker(1000 * time.Millisecond)
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
	}
	return "A", nil
}
func normalTask(ctx context.Context) (string, error) {
	t := time.NewTicker(3000 * time.Millisecond)
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
	}
	return "B", nil
}

func subTask01(ctx context.Context, wg *sync.WaitGroup, id string) {
	defer wg.Done()
	// task01の処理に500msかかると仮定している
	t := time.NewTicker(500 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	case <-t.C:
		t.Stop()
		fmt.Println(id)
	}
}

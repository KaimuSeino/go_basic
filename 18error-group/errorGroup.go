package errorgroup

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func ErrorGroup() {
	ctx, cancel := context.WithTimeout(context.Background(), 1600*time.Millisecond)
	defer cancel()
	// eg := new(errgroup.Group)
	eg, ctx := errgroup.WithContext(ctx)
	s := []string{"task1", "task2", "task3", "task4"}
	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(ctx, task)
		})
	}
	if err := eg.Wait(); err != nil { // エラーが来るまで待機
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("finish")
}

func doTask(ctx context.Context, task string) error {
	var t *time.Ticker
	switch task {
	case "task1":
		t = time.NewTicker(1000 * time.Millisecond)
	case "task2":
		t = time.NewTicker(1400 * time.Millisecond)
	default:
		t = time.NewTicker(2000 * time.Millisecond)
	}
	select {
	case <-ctx.Done():
		fmt.Printf("%v cancelled: %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()
		// if task == "fake1" || task == "fake2" {
		// 	return fmt.Errorf("%v process failed", task)
		// }
		fmt.Printf("task %v completed\n", task)
	}
	return nil
}

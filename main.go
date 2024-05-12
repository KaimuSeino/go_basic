package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// // mutex
	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// var i int
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	mu.Lock() // 他のgoroutineでiが操作できないようにロックをかける。
	// 	defer mu.Unlock()
	// 	// i++
	// 	i = 1
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	mu.Lock() // 他のgoroutineでiが操作できないようにロックをかける。
	// 	defer mu.Unlock()
	// 	// i++
	// 	i = 2
	// }()
	// wg.Wait()
	// fmt.Println(i)
	// // たまにi=0にi++が同時に起きて１が表示される。data raceが発生
	// // 確認方法 go run -race main.go
	// // データ競合はmutexで解決できる

	// // RWmutex
	// var wg sync.WaitGroup
	// var rwMu sync.RWMutex
	// var c int

	// wg.Add(4)
	// write(&rwMu, &wg, &c)
	// read(&rwMu, &wg, &c)
	// read(&rwMu, &wg, &c)
	// read(&rwMu, &wg, &c)

	// wg.Wait()
	// fmt.Println("finish")

	// atomic
	var wg sync.WaitGroup
	var c int64

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&c, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(c)
	fmt.Println("finish")

}
func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	// RWmutexを使うとread unlockを行わなくても、次のreadができる
	defer wg.Done()
	time.Sleep(10 * time.Microsecond)
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("read lock")
	fmt.Println(*c)
	time.Sleep(1 * time.Second)
	fmt.Println("read unlock")
}

func write(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("write lock")
	*c += 1
	time.Sleep(1 * time.Second)
	fmt.Println("write unlock")
}

package mutexatomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func MutexAtomic() {
	// mutex
	var wg sync.WaitGroup
	var mu sync.Mutex
	var i int
	wg.Add(2)
	go func() {
		defer wg.Done()
		mu.Lock() // 他のgoroutineでiが操作できないようにロックをかける。
		defer mu.Unlock()
		i++
	}()
	go func() {
		defer wg.Done()
		mu.Lock()         // ロックをかける。
		defer mu.Unlock() // i++の後にアンロック
		i++
	}()
	wg.Wait()
	fmt.Println(i)
	fmt.Println("finish")
	// たまにi=0にgoroutineのi++が同時に発生して1が表示されることがある。
	// データ競合はmutexで解決できる: go run -race file.name

	fmt.Println("########################")
	time.Sleep(3 * time.Second)
	// RWmutex
	var rwMu sync.RWMutex
	var c int

	wg.Add(4)
	go write(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)

	wg.Wait()
	fmt.Println("finish")

	fmt.Println("################")
	time.Sleep(3 * time.Second)

	// atomic
	// 排他的制御
	var v int64
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&v, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(v)
	fmt.Println("finish")
}

func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	// RWmutexを使うとread unlockを行う前でも、読み込みが可能になる
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
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
